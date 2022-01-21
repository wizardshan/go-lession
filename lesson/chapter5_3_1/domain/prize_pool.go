package domain

type PrizePool struct {
	Prizes []IPrize
}


func (dom *PrizePool) Lottery() (IPrize, error) {

	prize := dom.getPrizeByPercent(dom.Prizes)

	_, isThankPrize := prize.(*Thanks)
	if !isThankPrize {
		if err := prize.CheckLimit(); err != nil {
			return nil, err
		}
	}

	prize.Award()

	return prize, nil
}

func (dom *PrizePool) getPrizeByPercent(prizes []IPrize) IPrize {
	// 模拟随机抽一个
	return prizes[3]
}

