package ebs

import "github.com/cvila84/erpdump/pkg/pivot"

/*
2022 Delta with baseline/forecast: +1 SM in Praha agreed by Mauricio [R1R29750]
2022 Delta with baseline/forecast: +2 ppl in Noida agreed by Mauricio to compensate turn-overs [R1R29750]
2022 Delta with baseline/forecast: SII budget for OTA-BE & trainings approved by Nagy [R1R29750]
2022 Delta with baseline/forecast: Praha new infrastructure agreed by David/Guillaume [R1R29751]
2022 Delta with baseline/forecast: QA black raised 800$/month of AWS costs agreed by Mauricio [R1R29751]
2022 Delta with baseline/forecast: +1 ppl (13->14) for AOTA L3 agreed by Mauricio [R0S29752]
2022 Virtual BW used for: European digital wallet agreed by Samir Khlif (IBS) [R1R29753]
2022 Virtual BW used for: Private Network and xRIM aaS agreed by Mauricio [R1R29753]

2023 Delta with baseline/forecast: NFV improvements NOT APPROVED [R1R29750], inputs made by Daniel
*/

const (
	Workload   = "Workload"
	Travel     = "Travel & entertainment"
	Agency     = "Temps & agency costs"
	Recharge   = "Allocations & recharges mgmt"
	Employee   = "Employee-related"
	CenterCost = "DC Cost"
	OpCost     = "Operating costs"
	Fees       = "Professional fees"
	Facilities = "Facilities"
)

// ---

var otaDevProjects = []string{
	"R1R29750",
}

var otaCogsProjects = []string{
	"EBD21005",
	"EBD21015",
	"EBD22002",
	"EEF22017",
}

var otaOcosProjects = []string{
	"R0S29752",
}

var tacDevProjects = []string{
	"R1R30027",
}

var iotDevProjects = []string{
	"R1R30028",
}

var otaOtherDevProjects = []string{
	"R1R29751",
	"R1R29753",
	"R0R29754",
}

// ---

var otaPtfProjects = []string{
	"R1R29751",
}

var otaInnovationProjects = []string{
	"R1R29753",
}

var otaImprovementProjects = []string{
	"R0R29754",
}

var functionalOtherProjects = []string{
	"RDX0000A",
	"RDX0000S",
	"RDX0000T",
	"RDX000PT",
	"X0000T",
	"CWB10000",
	"CWT10000",
}

var functionalHolidaysProjects = []string{
	"RDX0000H",
	"CWH10000",
}

var projectGroups = func(prefixProject bool) pivot.Compute[string] {
	return func(elements []pivot.RawValue) (string, error) {
		project, ok := elements[0].(string)
		if !ok {
			return "", pivot.InvalidType(elements[0])
		}
		var prefix string
		if prefixProject {
			prefix = project + "-"
		}
		teamWorkload, ok := projectsWorkloadSplit[project]
		if ok {
			for k, v := range teamWorkload {
				for _, p := range v {
					if p == elements[1] && Workload == elements[2] {
						return prefix + k, nil
					}
				}
			}
		}
		otherCosts, ok := projectOtherCostsSplit[project]
		if ok {
			for k, v := range otherCosts {
				for _, p := range v {
					if p == elements[2] {
						return prefix + k, nil
					}
				}
			}
		}
		return prefix + "Unknown", nil
	}
}

type projectSplit map[string][]string

var projectsWorkloadSplit map[string]projectSplit

var projectOtherCostsSplit map[string]projectSplit

func initProjects(index int, verbose bool) {
	projectsWorkloadSplit = make(map[string]projectSplit)
	projectOtherCostsSplit = make(map[string]projectSplit)
	projectsWorkloadSplit["R1R29750"] = projectSplit{
		"Budget":       uniquePeople(verbose, index, cotaDevL3BudgetPeople),
		"Budget(vOTA)": uniquePeople(verbose, index, cotaPtfVMBudgetPeople),
		"Other(AOTA)":  uniquePeople(verbose, index, aotaL3OtherPeople, aotaDevOtherPeople),
		"Other":        uniquePeople(verbose, index, cotaDevL3OtherPeople),
		"Extension":    uniquePeople(verbose, index, ext29750MyosdTeamPeople, ext29750Tls13People, ext29750NewAppletsPeople, ext29750NgmMigrationPeople),
	}
	projectOtherCostsSplit["R1R29750"] = projectSplit{
		"Budget":    {Recharge},
		"Other":     {Employee, Travel},
		"Extension": {Agency},
	}
	projectsWorkloadSplit["R1R29751"] = projectSplit{
		"Budget":    uniquePeople(verbose, index, cotaPtfBudgetPeople),
		"Other":     uniquePeople(verbose, index, cotaPtfOtherPeople),
		"Extension": uniquePeople(verbose, index, ext29751OtaDemoTenantPeople),
	}
	projectOtherCostsSplit["R1R29751"] = projectSplit{
		"Budget": {OpCost, Fees, Facilities, CenterCost},
	}
	projectsWorkloadSplit["R0S29752"] = projectSplit{
		"Budget(COTA)": uniquePeople(verbose, index, cotaDevL3BudgetPeople, cotaPtfVMBudgetPeople),
		"Budget(AOTA)": uniquePeople(verbose, index, aotaL3BudgetPeople),
		"Other(COTA)":  uniquePeople(verbose, index, cotaDevL3OtherPeople),
		"Other(AOTA)":  uniquePeople(verbose, index, aotaL3OtherPeople),
	}
	projectsWorkloadSplit["R1R29753"] = projectSplit{
		"Budget": uniquePeople(verbose, index, innovationBudgetTransPeople),
		"Other":  uniquePeople(verbose, index, innovationOtherServerPeople, innovationOtherAppletPeople),
	}
	projectOtherCostsSplit["R1R29753"] = projectSplit{
		"Other": {CenterCost},
	}
	projectsWorkloadSplit["R0R29754"] = projectSplit{
		"Budget": uniquePeople(verbose, index, improvmentBudgetPeople),
		"Other":  uniquePeople(verbose, index, improvmentOtherPeople),
	}
	projectOtherCostsSplit["R0R29754"] = projectSplit{
		"Other": {Travel, CenterCost},
	}
	projectsWorkloadSplit["R0R29805"] = projectSplit{
		"Budget": uniquePeople(verbose, index, centralRDPeople),
	}
	projectOtherCostsSplit["R0R29805"] = projectSplit{
		"Other": {Agency},
	}
	projectsWorkloadSplit["R0T30005"] = projectSplit{
		"Budget": uniquePeople(verbose, index, transversalPeople),
	}
	projectsWorkloadSplit["R1R30027"] = projectSplit{
		"Budget":    uniquePeople(verbose, index, tacBudgetAppletPeople),
		"Other":     uniquePeople(verbose, index, tacOtherServerPeople),
		"Extension": uniquePeople(verbose, index, ext30027transatelActPeople),
	}
	projectsWorkloadSplit["R1R30028"] = projectSplit{
		"Budget": uniquePeople(verbose, index, iotBudgetServerPeople, iotBudgetAppletPeople, iotBudgetTransPeople),
		"Other":  uniquePeople(verbose, index, iotOtherPeople),
	}
}

func init() {
	initProjects(1, true)
}
