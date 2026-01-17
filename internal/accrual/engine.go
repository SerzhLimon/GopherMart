package accrual

import (
	"context"

	m "github.com/SerzhLimon/GopherMart/internal/models_accrual"
	repo "github.com/SerzhLimon/GopherMart/internal/repository_accrual"
	"github.com/sirupsen/logrus"
)

type AccrualEngine struct {
	repo              repo.Repository
	ordersToProcessCh chan m.Order
	stopCh            chan struct{}
}

func New(repo repo.Repository) *AccrualEngine {
	return &AccrualEngine{
		repo:              repo,
		ordersToProcessCh: make(chan m.Order),
		stopCh:            make(chan struct{}),
	}
}

func (p *AccrualEngine) RegisterOrderForProcessing(order m.Order) {
	for {
		select {
		case <-p.stopCh:
			logrus.Info("engine stopped")
			return
		case p.ordersToProcessCh <- order:
			logrus.Infof("%s registred for processing", order.ID)
			return
		}
	}
}

func (e *AccrualEngine) StartProcessing(ctx context.Context) {
	// newOrders, err := e.repo.GetOrdersForProcessing()
	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("cant fetch orders for processing")
	// }

	// for _, newOrder := range newOrders {
	// 	log.Debug().Int("order_id", newOrder.ID).Msg("start processing order")
	// 	go p.AddAccrualForOrder(ctx, newOrder)
	// }

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		// when context is canceled, we will signal all senders that they should stop
	// 		// sending orders to channel, because we won't process them anymore
	// 		close(p.stopCh)
	// 		log.Debug().Msg("stop receiving orders to process")
	// 		return
	// 	case orderToProcess := <-p.ordersToProcessCh:
	// 		log.Debug().Int("order_id", orderToProcess.ID).Msg("received order to process")
	// 		go p.AddAccrualForOrder(ctx, orderToProcess)

	// 	}
	// }
}
