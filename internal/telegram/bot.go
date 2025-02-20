package telegram

import (
	"TelegramBot/internal/telegram/fsm"
	"TelegramBot/internal/telegram/models/request"
	"TelegramBot/internal/telegram/models/response"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const defaultTimeout = time.Minute
const defaultWorkers = 10

type Bot struct {
	token       string
	middlewares []Middleware
	router      *Router
	apiUrl      string
	lastUpdate  int64
	poolTimeout time.Duration
	httpClient  *http.Client
	updates     chan response.Update
	workers     int
	fsm         *fsm.MultiTypeMemoryFSM
}

func NewBot(token string, router *Router) *Bot {
	return &Bot{
		token:       token,
		apiUrl:      "https://api.telegram.org/bot" + token,
		router:      router,
		poolTimeout: defaultTimeout,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		updates: make(chan response.Update),
		workers: defaultWorkers,
		fsm:     fsm.NewMultiTypeFSM(),
	}
}

type Middleware interface{}

func (b *Bot) Start() {
	var wg sync.WaitGroup
	wg.Add(b.workers)
	for i := 0; i < b.workers; i++ {
		go b.processUpdate(&wg)
	}
	for {
		requestUpdate := request.Update{
			Timeout: int(b.poolTimeout.Seconds()),
			Offset:  atomic.LoadInt64(&b.lastUpdate) + 1,
		}
		updates, err := b.getUpdates(requestUpdate)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, update := range updates {
			b.updates <- update
			atomic.StoreInt64(&b.lastUpdate, update.UpdateId)

		}
	}
}

func (b *Bot) processUpdate(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case update := <-b.updates:
			b.router.mu.Lock()
			if targetHandlers, ok := b.router.handlers[update.Type()]; ok {
				for _, h := range targetHandlers {
					if !h.Filter(NewContext(b, &update)) {
						continue
					}
					switch update.Type() {
					case response.MessageHandler:
						h.HandleMessage(NewContextMessage(b, &update))
					case response.CallbackQueryHandler:
						h.HandleCallbackQuery(NewContextCallbackQuery(b, &update))
					}
					break
				}
			}
			b.router.mu.Unlock()
		}
	}
}

func (b *Bot) FSM() *fsm.MultiTypeMemoryFSM {
	return b.fsm
}
