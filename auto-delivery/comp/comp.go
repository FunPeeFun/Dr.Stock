package comp

import (
	"../buffer"
	"encoding/json"
	"fmt"
)

type AnalRes struct {
	// 角色比例：后卫、中锋、前锋
	RolePercent map[string] float32
	// 设置的总比例(占仓位的比例，而不是总金额的比例)
	// 如果超过100%，则认为超出了预期
	SetPercentSum float32
	// 已投资金额
	InvestedAmount float32
	// 计划投资金额
	ToInvestAmount float32
	// 加仓计划
	InvestPlans map[string] *InvestPlan
}

type InvestPlan struct {
	Name string
	// 加仓速率，元/周
	InvestPerWeek float32
	// 总加仓金额，+表示加仓，-表示减仓
	ToInvestAmount float32
}

var AnalysisResult = AnalRes{RolePercent: make(map[string] float32),
	SetPercentSum: 0,
	InvestedAmount: 0,
	InvestPlans: make(map[string] *InvestPlan)}

func Analysis() {
	buffer.StockConfigLock.Lock()
	stocks := buffer.StockConfig.Stocks
	AnalysisResult.ToInvestAmount = buffer.StockConfig.TotalAmount * buffer.StockConfig.SetTotalPercent / 100.0
	for stockNum, stock := range stocks {
		AnalysisResult.SetPercentSum += stock.SetPercent
		AnalysisResult.RolePercent[stock.Role] += stock.SetPercent
		AnalysisResult.InvestedAmount += stock.Amount
		AnalysisResult.InvestPlans[stockNum] = &InvestPlan{"", 0 , 0}
			AnalysisResult.InvestPlans[stockNum].Name = stock.Name
		AnalysisResult.InvestPlans[stockNum].ToInvestAmount = AnalysisResult.ToInvestAmount *
			stock.SetPercent / 100.0 -
			stock.Amount
		AnalysisResult.InvestPlans[stockNum].InvestPerWeek = AnalysisResult.InvestPlans[stockNum].ToInvestAmount /
			float32(buffer.StockConfig.FixedInvestWeek)
	}
	buffer.StockConfigLock.Unlock()
	res, err := json.MarshalIndent(AnalysisResult, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", res)
}
