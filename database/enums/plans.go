package enums

type Plan string

const (
	PlanTierNone         Plan = "none"
	PlanTierProfessional Plan = "PRO"
	PlanTierPremium      Plan = "PRM"
)

func (p Plan) IsValid() bool {
	return p == PlanTierNone || p == PlanTierPremium || p == PlanTierProfessional
}
