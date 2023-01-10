package ebs

import (
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
)

/*
Delta with baseline/forecast: +1 SM in Praha agreed by Mauricio [R1R29750]
Delta with baseline/forecast: +2 ppl in Noida agreed by Mauricio to compensate turn-overs [R1R29750]
Delta with baseline/forecast: Praha new infrastructure agreed by ? [R1R29751]
Delta with baseline/forecast: QA black raised 800$/month of AWS costs agreed by ? [R1R29751]
Delta with baseline/forecast: +1 ppl (13->14) for AOTA L3 agreed by Mauricio [R0S29752]
*/

// Delta with baseline/forecast: TLS 1.3 development (93md on waf+psk-provider) approved by Nagy [R1R29750]
var ext29750Tls13People = [][]string{
	{"Bretagne,Eric", "Bretagne Eric"},
	{"", "Oumohand Rachid"},
	{"", "Kumar Singh Shishir"},
}

// Delta with baseline/forecast: NGM-Datadog migration agreed by Mauricio [R1R29750]
var ext29750NgmMigrationPeople = [][]string{
	{"", "Toschi Guilherme"},
}

// Delta with baseline/forecast: New refresh applet, MIV2 study for NB-IOT, Applet study to transform 4G card into 5G agreed by Nagy [R1R29750]
var ext29750NewAppletsPeople = [][]string{
	{"", "Abao Michael Carlo"},
	{"", "Giva Joana Marie"},
	{"", "Marquez Justin"},
}

// Delta with baseline/forecast: MyOSD team contribution [R1R29750]
var ext29750MyosdTeamPeople = [][]string{
	{"", "Navarrete Perez Hector Luis"},
}

// Delta with baseline/forecast: OTA demo tenant approved by Jérome Voyer [R1R29751]
var ext29751OtaDemoTenantPeople = [][]string{
	{"", "Virmani Karan"},
}

// Delta with baseline/forecast: European digital wallet agreed by Samir Khlif (IBS) [R1R29753]
var ext29753DigitalWalletPeople = [][]string{
	{"Gattone,Alain", "Gattone Alain"},
	{"Berard,Xavier", "Berard Xavier"},
}

// Delta with baseline/forecast: Private Network and xRIM aaS agreed by Mauricio [R1R29753]
var ext29753privateNetworkPeople = [][]string{
	{"", "Prigent Francois"},
}

// Delta with baseline/forecast: Transatel DP+ flow activation for 3k€ [R1R30027]
var ext30027transatelActPeople = [][]string{
	{"", "Ayasse Jerome"},
}

// R0R29805 Central R&D
// 366k€
var centralRDPeople = [][]string{
	{"", "Chantre Thierry"},
	{"", "Detcheverry Frank"},
	{"", "Jaramilla Michael Christian"},
	{"", "Kehyayan Stephane"},
	{"", "Lacouture Dominique"},
	{"", "Lambert Patrick"},
	{"", "Leal Rainier"},
	{"", "Maille Caroline"},
	{"", "Mateo Vladimir Lennard"},
	{"", "Royeca Zigmund Zer"},
	{"", "Salem Butch Ellen Grace"},
	{"", "Scheerders Bruno"},
	{"", "Tacussel Philippe"},
	{"", "Theil Fabienne"},
}

// R0T30005 Transversal
// 478k€
var transversalPeople = [][]string{
	{"", "Almanza Roman Ines Atenea"},
	{"Bretagne,Eric", "Bretagne Eric"},
	{"", "Dubuc Laurent"},
	{"", "Garduno Sanchez Victor Hugo"},
	{"", "Gulyani Sahil"},
	{"", "Huerta martinez Jesus"},
	{"", "Kumar Pramod"},
	{"Kumar,Vivek", "Kumar Vivek"},
	{"", "Loos Michel"},
	{"", "Manchanda Jatin"},
	{"", "Patel Kamal"},
	{"", "Prakash Virender"},
	{"", "Prigent Francois"},
	{"", "Qualizza Michele"},
	{"", "Singh Alok"},
	{"", "Singh Chandan Kumar"},
	{"", "Singh Gaurav"},
	{"", "Tcherniack Laurent"},
	{"", "Tovar Jonathan Josué"},
	{"", "Yadav Sanjeet"},
}

// R0R29754 Improvement
// 2ppl/12m/100%
// Kriss: 3420h / 168435€
// Vertical: 2992h / 147356€
var improvmentBudgetPeople = [][]string{
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"},
	{"Oremus,Tomas", "Oremus Tomas"},
}
var improvmentOtherPeople = [][]string{
	{"ATMOPAWIRO,ALSASIAN", "ATMOPAWIRO ALSASIAN"},
	{"Berard,Xavier", "Berard Xavier"},
	{"Cerny,Jaroslav", "Cerny Jaroslav"},
	{"Gattone,Alain", "Gattone Alain"},
	{"", "Halagunde Vinayak Punja"},
	{"", "Marciniak Mateusz"},
	{"", "Rajput Yashwant Singh"},
	{"", "Rana Ritika"},
	{"", "Ververis Konstantinos"},
	{"", "Vikas Vikas"},
}

// R1R29753 Innovation
// 2.5ppl/12m/100%
// Kriss: 3971h / 413977€
// Vertical: 3740h / 389895€
var innovationBudgetPeople = [][]string{
	{"", "Abao Michael Carlo"},
	{"", "ABAO MICHAEL CARLO"},
	{"Bretagne,Eric", "Bretagne Eric"},
	{"", "Castano Esteban"},
	{"Cerny,Jaroslav", "Cerny Jaroslav"},
	{"Deepak,Deepak", "Deepak Deepak"},
	{"Dhondiyal,Rituraj", "Dhondiyal Rituraj"},
	{"", "Eleserio Ederlyn"},
	{"", "Freimonas Romas"},
	{"", "Giva Joana Marie"},
	{"", "Kobr Dan"}, // KMT
	{"Kumar,Anshuman", "Kumar Anshuman"},
	{"Lachowicz,Daniel", "Lachowicz Daniel"},
	{"", "LAI SER WEI"},
	{"", "Marquez Justin"},
	{"", "Miani Alberto"}, // KMT
	{"", "ONG WILSON LEE"},
	{"", "Sharma Gaurav"},
	{"", "Virmani Karan"},
}
var innovationOtherPeople = [][]string{}

// R1R29750 AOTA Dev
// No budget
var aotaDevBudgetPeople = [][]string{}
var aotaDevOtherPeople = [][]string{
	{"Agrawal,Somya", "Agrawal Somya"},
	{"Arora,Sheffali", "Arora Sheffali"},
	{"Choubisa,Vidhi", "Choubisa Vidhi"},
}

// R0S29752 AOTA L3
// 13ppl/12m/90%
// Kriss: 17609h / 867227€
// Vertical: 17503h / 862033€
var aotaL3BudgetPeople = [][]string{
	{"Agrawal,Somya", "Agrawal Somya"},
	{"Arora,Sheffali", "Arora Sheffali"},
	{"Bhatia,Akhil", "Bhatia Akhil"},
	{"Choubisa,Vidhi", "Choubisa Vidhi"},
	{"Gupta,Nainsy", "Gupta Nainsy"},
	{"Kumar Verma,Ayush", "Kumar Verma Ayush"},
	{"Kumar,Sachin", "Kumar Sachin"},
	{"Kushwaha,Amarjit", "Kushwaha Amarjit"},
	{"Moudgil,Rhea", "Moudgil Rhea"},
	{"Patil,Ramesh", "Patil Ramesh"},
	{"RAJU,RAMESH", "RAJU RAMESH"},
	{"Shameem,Bilal", "Shameem Bilal"},
	{"Sharma,Naveen", "Sharma Naveen"},
}
var aotaL3OtherPeople = [][]string{
	{"Devanjali,Devanjali", "Devanjali Devanjali"},
	{"Dubey,Parul", "Dubey Parul"},
	{"GOYAL,JAIDEV", "GOYAL JAIDEV"},
	{"Gupta,Swati", "Gupta Swati"},
	{"", "Kalra Prashant"},
	{"KUMAR,Rahul", "KUMAR Rahul"},
	{"Martinez Carino,Conrado", "Martinez Carino Conrado"},
	{"", "Ocampo Gonzalez Francisco Javier"},
	{"Perez Cuellar,Julio Cesar", "Perez Cuellar Julio Cesar"},
	{"SAHADEVAN,SANU", "SAHADEVAN SANU"},
	{"Srivastava,Sahil", "Srivastava Sahil"},
}

// R1R29751 COTA Platform SaaS
// (PT/FR) 3ppl/12m/100%
// (PT/IN) 5ppl/12m/100%
// Kriss: 12211h / 601402€
// Vertical: 4488h / 221034€
// R1R29751 COTA Platform NFV
// (PT/FR) 1ppl/12m/100%
// (PT/CZ) 4ppl/6m/50%
// Kriss: 3297h / 162353€
// Vertical: 1496h / 73678€
var cotaPtfBudgetPeople = [][]string{
	// SaaS
	{"Deepak,Deepak", "Deepak Deepak"},
	{"Dhondiyal,Rituraj", "Dhondiyal Rituraj"},
	{"LAMBA,PREETIKA", "LAMBA PREETIKA"},
	{"Levacher,Frederic", "Levacher Frederic"},
	{"Ragot,Emmanuel", "Ragot Emmanuel"},
	{"Roucheton,Jerome", "Roucheton Jerome"},
	{"Singh,Pradeep", "Singh Pradeep"},
	{"Singh,Yash Pal", "Singh Yash Pal"},
	// Strategic
	{"Chiaramello,Daniel", "Chiaramello Daniel"},
	{"Dokladal,Jakub", "Dokladal Jakub"},
	{"Lachowicz,Daniel", "Lachowicz Daniel"},
	{"Przytarski,Bartlomiej", "Przytarski Bartlomiej"},
	{"Sedlacek,Ondrej", "Sedlacek Ondrej"},
}
var cotaPtfOtherPeople = [][]string{
	{"", "Aguilera Palomino Diego"}, // support from deployment team
	{"", "Alarcon galvez Fernando"}, // support from deployment team
	//{"Bories,Clement", "Bories Clement"},
	{"BHATNAGAR,AAKARSH", "BHATNAGAR AAKARSH"}, // newcomer to replace budgeted resource
	{"Bretagne,Eric", "Bretagne Eric"},         // KMT
	//{"Cabrera,Marcos", "Cabrera Marcos"},
	{"", "Castano Esteban"},                                // support from L2
	{"Cerny,Jaroslav", "Cerny Jaroslav"},                   // KMT
	{"Delgado martinez,Alvaro", "Delgado martinez Alvaro"}, // KMT
	{"Fioux,Sebastien", "Fioux Sebastien"},                 // SaaS platform team SM
	{"", "Freimonas Romas"},                                // KMT
	{"", "Galindo Gomez Jorge"},                            // ?
	{"Gukanti,Sandeep", "Gukanti Sandeep"},                 // support from deployment team
	{"", "Jones Terence"},                                  // support from DB expert team
	{"Kumar,Anshuman", "Kumar Anshuman"},                   // support from deployment team
	{"Letolle,Nicolas", "Letolle Nicolas"},                 // support from deployment team
	{"LIM, Mr BERNARD KENNETH", "LIM BERNARD KENNETH"},     // support from DB expert team
	{"", "Nandiraju Pavan Kumar"},                          // ?
	{"", "Perez Lagunas Daniela"},                          // support from MyOSD team
	{"Schammass,Alexandre", "Schammass Alexandre"},         // support from deployment team
	{"Sharma,Aditya", "Sharma Aditya"},                     // newcomer to replace budgeted resource
	{"", "Sharma Gaurav"},                                  // deployment/platform team manager
	{"", "Valette Karine"},                                 // support from DB expert team
}

// R1R29750 COTA Dev
// (Chiefs/FR) 3ppl/12m/100%
// (PDA/CZ) 1ppl/12m/100%
// (FT/FR) 1ppl/12m/100%
// (FT/CZ) 13ppl/12m/90%
// (FT/IN) 40ppl/12m/90%
// (VM/CZ) 5ppl/6m/50%
// Kriss: 83816h / 4127913€
// Vertical: 80709h / 3974928x€
// R0S29752 COTA L3
// (FT/CZ) 13ppl/12m/10%
// (FT/IN) 40ppl/12m/10%
// Kriss: 8180h / 402855€
// Vertical: 7929h / 390493€
var cotaDevL3BudgetPeople = [][]string{
	// Chiefs/FR
	{"Cabagno,Anne", "Cabagno Anne"},
	{"Fioux,Sebastien", "Fioux Sebastien"},
	{"Tessier,Alexandra", "Tessier Alexandra"},
	// PDA/CZ
	{"Barilly,Adrien", "Barilly Adrien"},
	// FT/FR
	{"Coste,Florent", "Coste Florent"},
	// FT/CZ
	{"Delgado martinez,Alvaro", "Delgado martinez Alvaro"},
	{"Gatica Peralta,Elia Azucena", "Gatica Peralta Elia Azucena"},
	{"Gorokhov,Nikita", "Gorokhov Nikita"},
	{"Gorysz,Lukasz", "Gorysz Lukasz"},
	{"Hamid,Juba", "Hamid Juba"},
	{"Hlavacek,Ludek", "Hlavacek Ludek"},
	{"Jiricek,Libor", "Jiricek Libor"},
	{"Kalita,Victoria", "Kalita Victoria"},
	{"Kostohryz,Jan", "Kostohryz Jan"},
	{"Manasek,Radek", "Manasek Radek"},
	{"Nguyen Gia Can,Cyril", "Nguyen Gia Can Cyril"},
	{"Norko,Veronika", "Norko Veronika"},
	{"Penha,Bruno", "Penha Bruno"},
	{"Pomorski,Marcin", "Pomorski Marcin"},
	{"Sramek,Vaclav", "Sramek Vaclav"},
	{"Uzun,Eraslan", "Uzun Eraslan"},
	{"Vondracek,Martin", "Vondracek Martin"},
	// FT/IN
	{"Agrawal,Deepak", "Agrawal Deepak"},
	{"Agrawal,Ritika", "Agrawal Ritika"},
	{"Ali,Riyasat", "Ali Riyasat"},
	{"BANSAL,ANKIT", "BANSAL ANKIT"},
	{"Bhardwaj,Abhishek", "Bhardwaj Abhishek"},
	{"Bhatia,Akhil", "Bhatia Akhil"},
	{"Bhatnagar,Manas", "Bhatnagar Manas"},
	{"Dhondiyal,Rituraj", "Dhondiyal Rituraj"},
	{"Chauhan,Yashvant Singh", "Chauhan Yashvant Singh"},
	{"Gupta,Ankur", "Gupta Ankur"},
	{"Gupta,Anshul", "Gupta Anshul"},
	{"Harshadbhai,Patel Kaushikkumar", "Harshadbhai Patel Kaushikkumar"},
	{"Jain,Himanshu", ""}, // not present in finance
	{"Jain,Shubham", "Jain Shubham"},
	{"Jha,Ashwani Kumar", "Jha Ashwani Kumar"},
	{"Kalani,Anukriti", "Kalani Anukriti"},
	{"KANSAL,MANIKA", ""}, // not present in finance
	{"Khan,Akram Raza", "Khan Akram Raza"},
	{"Khan,Momin", "Khan Momin"},
	{"Khan,Wasim", "Khan Wasim"},
	{"Kumar,Amrish", "Kumar Amrish"},
	{"Kumar,Chandan", "Kumar Chandan"},
	{"Kumar,Gaurav", "Kumar Gaurav"},
	{"KUMAR,Narendra", "KUMAR Narendra"},
	{"Kumar,Satyam", "Kumar Satyam"},
	{"kumar,Saurabh", "kumar Saurabh"},
	{"Kumar,Sujit", "Kumar Sujit"},
	{"Kumar,Vikas", "Kumar Vikas"},
	{"Kumar,Vivek", "Kumar Vivek"},
	{"PANDEY,Priyesh", "PANDEY Priyesh"},
	{"Pandey,Renu", "Pandey Renu"},
	{"Pethia,Abhishek", "Pethia Abhishek"},
	{"Priyanka,Kumari", "Priyanka Kumari"},
	{"Rajwaar,Shweta", "Rajwaar Shweta"},
	{"Rawat,Yoginder", "Rawat Yoginder"},
	{"Rupera,Bhumika", "Rupera Bhumika"},
	{"SHARMA,ADITI", "SHARMA ADITI"},
	{"Sharma,Ashwani Kumar", "Sharma Ashwani Kumar"},
	{"Sharma,Himanshu", "Sharma Himanshu"},
	{"Sharma,Prateek", "Sharma Prateek"},
	{"Sharma,Roshan Lal", "Sharma Roshan Lal"},
	{"Singh,Akshay", "Singh Akshay"},
	{"Singh,Ashutosh", "Singh Ashutosh"},
	{"Singh,Bheem", "Singh Bheem"},
	{"Singh,Devyani", "Singh Devyani"},
	{"Singh,Gurvinder", "Singh Gurvinder"},
	{"Singh,Hameer", "Singh Hameer"},
	{"SINGH,VIJAYLAXMI", "SINGH VIJAYLAXMI"},
	{"Singhal,Shivank", "Singhal Shivank"},
	{"Sirvya,Anshul", "Sirvya Anshul"},
	{"Vats,Vishant Kumar", ""}, // not present in finance
	{"VOHRA,Mitali", "VOHRA Mitali"},
	// VM/CZ
	{"Dokladal,Jakub", "Dokladal Jakub"},
	{"", "Fedai Artem"},
	{"Lachowicz,Daniel", "Lachowicz Daniel"},
	{"Przytarski,Bartlomiej", "Przytarski Bartlomiej"},
	{"Sedlacek,Ondrej", "Sedlacek Ondrej"},
}
var cotaDevL3OtherPeople = [][]string{
	{"Chiaramello,Daniel", "Chiaramello Daniel"},                                 // support from platform team
	{"Deepak,Deepak", "Deepak Deepak"},                                           // support from platform team
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"}, // support from SSC
}

// R1R30027 TAC Dev
// 2ppl/12m/25%
// Kriss: 847h / 40781€
// Vertical: 748h / 35998€
var tacBudgetPeople = [][]string{
	// Applet
	{"", "Abao Michael Carlo"},
	{"", "ONG WILSON LEE"},
	// Server
}
var tacOtherPeople = [][]string{
	// Applet
	{"", "Eleserio Ederlyn"},
	{"", "Shamsudin Nurrasyidah"},
	// Server
	{"", "Dumitrescu Florin"},      // support from ODC
	{"Gupta,Ankur", "Gupta Ankur"}, // server development was not in budget
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"}, // support from SSC
	{"", "Myslivets Alexey"},               // support from ODC
	{"", "Shevnin Ignat"},                  // support from ODC
	{"Singh,Gurvinder", "Singh Gurvinder"}, // server development was not in budget
	{"Singhal,Shivank", "Singhal Shivank"}, // server development was not in budget
}

// R1R30028 IOT Dev
// (Experts) 2ppl/12m/25%
// (FT/IN) 1ppl/12m/100%
// (Applet) 1ppl/12m/25%
// Kriss: 2734h / 179302€
// Vertical: 2618h / 169656€
var iotBudgetPeople = [][]string{
	// Applet
	{"", "LAI SER WEI"},
	// Server
	{"ATMOPAWIRO,ALSASIAN", "ATMOPAWIRO ALSASIAN"},
	{"Gattone,Alain", "Gattone Alain"},
	{"Kumar,Vishesh", "Kumar Vishesh"},
}
var iotOtherPeople = [][]string{
	// Applet
	{"", "Abao Michael Carlo"},
	{"", "Eleserio Ederlyn"},
	{"", "Espinosa Alen"},
	// Server
	{"Kumar,Vivek", "Kumar Vivek"}, // server development was not in budget
	{"", "Prigent Francois"},
	{"", "Yadav Sanjeet"},
}

var managers = []string{
	"Pereira Carrari, Mr Mauricio",
	"Letolle,Nicolas",
	"Vila,Christophe",
	"Kumar,Vikash",
	"Franco Mora,Richard Miguel",
	"Fesquet,Sebastien",
	"Pethia,Abhishek",
	"Khan,Akram Raza",
	"Mehta,Ashish",
	"KUMAR,Narendra",
	"Gupta,Anshul",
	"Kumar,Vivek",
}

type projectTeam struct {
	budget    []string
	extension []string
	other     []string
}

var projectsTeam map[string]projectTeam

var projectGroups = func(prefixProject bool) pivot.Compute[string] {
	return func(elements []string) string {
		var prefix string
		if prefixProject {
			prefix = elements[0] + "-"
		}
		team, ok := projectsTeam[elements[0]]
		if ok {
			for _, p := range team.budget {
				if p == elements[1] {
					return prefix + "Budget"
				}
			}
			for _, p := range team.extension {
				if p == elements[1] {
					return prefix + "Ext"
				}
			}
			for _, p := range team.other {
				if p == elements[1] {
					return prefix + "Other"
				}
			}
		}
		return prefix + "Unknown"
	}
}

func uniquePeople(index int, peopleLists ...[][]string) []string {
	var result []string
	for _, l1 := range peopleLists {
		for _, l2 := range l1 {
			if len(l2[index]) > 0 {
				present := false
				for _, p := range result {
					if l2[index] == p {
						present = true
						fmt.Printf("WARNING: duplicated people detected: %q\n", p)
					}
				}
				if !present {
					result = append(result, l2[index])
				}
			}
		}
	}
	return result
}

func init() {
	projectsTeam = make(map[string]projectTeam)
	projectsTeam["R1R29750"] = projectTeam{
		budget:    uniquePeople(1, aotaDevBudgetPeople, cotaDevL3BudgetPeople),
		extension: uniquePeople(1, ext29750MyosdTeamPeople, ext29750Tls13People, ext29750NewAppletsPeople, ext29750NgmMigrationPeople),
		other:     uniquePeople(1, aotaDevOtherPeople, cotaDevL3OtherPeople),
	}
	projectsTeam["R1R29751"] = projectTeam{
		budget:    uniquePeople(1, cotaPtfBudgetPeople),
		extension: uniquePeople(1, ext29751OtaDemoTenantPeople),
		other:     uniquePeople(1, cotaPtfOtherPeople),
	}
	projectsTeam["R0S29752"] = projectTeam{
		budget: uniquePeople(1, aotaL3BudgetPeople, cotaDevL3BudgetPeople),
		other:  uniquePeople(1, aotaL3OtherPeople, cotaDevL3OtherPeople),
	}
	projectsTeam["R1R29753"] = projectTeam{
		budget:    uniquePeople(1, innovationBudgetPeople),
		extension: uniquePeople(1, ext29753DigitalWalletPeople, ext29753privateNetworkPeople),
		other:     uniquePeople(1, innovationOtherPeople),
	}
	projectsTeam["R0R29754"] = projectTeam{
		budget: uniquePeople(1, improvmentBudgetPeople),
		other:  uniquePeople(1, improvmentOtherPeople),
	}
	projectsTeam["R0R29805"] = projectTeam{
		budget: uniquePeople(1, centralRDPeople),
	}
	projectsTeam["R0T30005"] = projectTeam{
		budget: uniquePeople(1, transversalPeople),
	}
	projectsTeam["R1R30027"] = projectTeam{
		budget:    uniquePeople(1, tacBudgetPeople),
		extension: uniquePeople(1, ext30027transatelActPeople),
		other:     uniquePeople(1, tacOtherPeople),
	}
	projectsTeam["R1R30028"] = projectTeam{
		budget: uniquePeople(1, iotBudgetPeople),
		other:  uniquePeople(1, iotOtherPeople),
	}
}
