package domain

type PrizePool struct {
	Prizes []IPrize
}


func (dom *PrizePool) Lottery() (IPrize, error) {

	prize := dom.getPrizeByPercent(dom.Prizes)

	p, ok := prize.(IPrizeLimiter)
	if ok {
		if err := p.CheckLimit(); err != nil {
			return nil, err
		}
	}

	prize.Award()

	return prize, nil
}

func (dom *PrizePool) getPrizeByPercent(prizes []IPrize) IPrize {
	// 模拟随机抽一个
	return prizes[0]
}

