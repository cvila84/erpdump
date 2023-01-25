package ebs

import (
	"fmt"
	"github.com/cvila84/erpdump/pkg/pivot"
)

// --- R0R29805 Central R&D
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

// --- R0T30005 Transversal
var transversalPeople = [][]string{
	{"Almanza Roman,Ines Atenea", "Almanza Roman Ines Atenea"},
	{"Bretagne,Eric", "Bretagne Eric"},
	{"Dubuc,Laurent", "Dubuc Laurent"},
	{"Garduno Sanchez,Victor Hugo", "Garduno Sanchez Victor Hugo"},
	{"Gulyani,Sahil", "Gulyani Sahil"},
	{"Huerta martinez,Jesus", "Huerta martinez Jesus"},
	{"Kumar,Pramod", "Kumar Pramod"},
	{"Kumar,Vivek", "Kumar Vivek"},
	{"Loos,Michel", "Loos Michel"},
	{"Manchanda,Jatin", "Manchanda Jatin"},
	{"Patel,Kamal", "Patel Kamal"},
	{"Prakash,Virender", "Prakash Virender"},
	{"Prigent,Francois", "Prigent Francois"},
	{"Qualizza,Michele", "Qualizza Michele"},
	{"Singh,Alok", "Singh Alok"},
	{"Singh,Chandan Kumar", "Singh Chandan Kumar"},
	{"Singh,Gaurav", "Singh Gaurav"},
	{"Tcherniack,Laurent", "Tcherniack Laurent"},
	{"Tovar,Jonathan Josué", "Tovar Jonathan Josué"},
	{"Yadav,Sanjeet", "Yadav Sanjeet"},
}

// --- R0R29754 Improvement
var improvmentBudgetPeople = [][]string{
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"},
	{"Oremus,Tomas", "Oremus Tomas"},
}
var improvmentOtherPeople = [][]string{
	{"ATMOPAWIRO,ALSASIAN", "ATMOPAWIRO ALSASIAN"},
	{"Berard,Xavier", "Berard Xavier"},
	{"Cerny,Jaroslav", "Cerny Jaroslav"},
	{"Gattone,Alain", "Gattone Alain"},
	{"Halagunde,Vinayak Punja", "Halagunde Vinayak Punja"},
	{"Marciniak,Mateusz", "Marciniak Mateusz"},
	{"Przytarski,Bartlomiej", "Przytarski Bartlomiej"},
	{"Rajput,Vikash Kumar", "Rajput Yashwant Singh"},
	{"Rana,Ritika", "Rana Ritika"},
	{"Ververis,Konstantinos", "Ververis Konstantinos"},
	{"Vikas,Vikas", "Vikas Vikas"},
}

// --- R0S29752 AOTA L3
var aotaL3BudgetPeople = [][]string{
	{"Agrawal,Somya", "Agrawal Somya"},
	{"Arora,Sheffali", "Arora Sheffali"},
	{"Choubisa,Vidhi", "Choubisa Vidhi"},
	{"Devanjali,Devanjali", "Devanjali Devanjali"}, // not named in budget but part of Ramesh team
	{"Dubey,Parul", "Dubey Parul"},                 // not named in budget but part of Ramesh team
	{"GOYAL,JAIDEV", "GOYAL JAIDEV"},               // not named in budget but part of Nainsy team
	{"Gupta,Nainsy", "Gupta Nainsy"},
	{"Gupta,Swati", "Gupta Swati"},       // not named in budget but part of Nainsy team
	{"Kalra,Prashant", "Kalra Prashant"}, // not named in budget but was probably part of L3 team
	{"KUMAR,Rahul", "KUMAR Rahul"},       // not named in budget but was probably part of L3 team
	{"Kumar,Sachin", "Kumar Sachin"},
	{"Kumar Verma,Ayush", "Kumar Verma Ayush"},
	{"Kushwaha,Amarjit", "Kushwaha Amarjit"},
	{"Moudgil,Rhea", "Moudgil Rhea"},
	{"Patil,Ramesh", "Patil Ramesh"},
	{"RAJU,RAMESH", "RAJU RAMESH"},
	{"SAHADEVAN,SANU", "SAHADEVAN SANU"}, // not named in budget but part of Nainsy team
	{"Shameem,Bilal", "Shameem Bilal"},
	{"Sharma,Naveen", "Sharma Naveen"},
	{"Srivastava,Sahil", "Srivastava Sahil"}, // not named in budget but part of Ramesh team
}
var aotaL3OtherPeople = [][]string{
	{"Berard,Xavier", "Berard Xavier"},                                       // ?
	{"Gattone,Alain", "Gattone Alain"},                                       // ?
	{"LIM, Mr BERNARD KENNETH", "LIM BERNARD KENNETH"},                       // ?
	{"Martinez Carino,Conrado", "Martinez Carino Conrado"},                   // ?
	{"Ocampo Gonzalez,Francisco Javier", "Ocampo Gonzalez Francisco Javier"}, // ?
	{"Perez Cuellar,Julio Cesar", "Perez Cuellar Julio Cesar"},               // ?
	{"Rajput,Vikash Kumar", "Rajput Yashwant Singh"},                         // ?
}

// --- R1R29750 COTA Dev
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
}
var cotaDevL3OtherPeople = [][]string{
	{"Chiaramello,Daniel", "Chiaramello Daniel"},                                 // support from platform team
	{"Deepak,Deepak", "Deepak Deepak"},                                           // support from platform team
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"}, // support from SSC
}
var cotaPtfVMBudgetPeople = [][]string{
	// VM/CZ
	{"Dokladal,Jakub", "Dokladal Jakub"},
	{"Fedai,Artem", "Fedai Artem"},
	{"Lachowicz,Daniel", "Lachowicz Daniel"},
	{"Przytarski,Bartlomiej", "Przytarski Bartlomiej"},
	{"Sedlacek,Ondrej", "Sedlacek Ondrej"},
}
var aotaDevOtherPeople = [][]string{
	{"Agrawal,Somya", "Agrawal Somya"},
	{"Arora,Sheffali", "Arora Sheffali"},
	{"Choubisa,Vidhi", "Choubisa Vidhi"},
}

// Delta with baseline/forecast: TLS 1.3 development (93md on waf+psk-provider) approved by Nagy [R1R29750]
var ext29750Tls13People = [][]string{
	{"Bretagne,Eric", "Bretagne Eric"},
	{"Kumar Singh,Shishir", "Kumar Singh Shishir"},
	{"Oumohand,Rachid", "Oumohand Rachid"},
}

// Delta with baseline/forecast: NGM-Datadog migration agreed by Mauricio [R1R29750]
var ext29750NgmMigrationPeople = [][]string{
	{"Toschi,Guilherme", "Toschi Guilherme"}, // previous manager Aroua,Adnen
}

// Delta with baseline/forecast: New refresh applet, MIV2 study for NB-IOT, Applet study to transform 4G card into 5G agreed by Nagy [R1R29750]
var ext29750NewAppletsPeople = [][]string{
	{"Abao,Michael Carlo", "Abao Michael Carlo"},
	{"Giva,Joana Marie", "Giva Joana Marie"},
	{"Marquez,Justin", "Marquez Justin"},
}

// Delta with baseline/forecast: MyOSD team contribution [R1R29750]
var ext29750MyosdTeamPeople = [][]string{
	{"Navarrete Perez,Hector Luis", "Navarrete Perez Hector Luis"},
}

// --- R1R29751 COTA Platform SaaS
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
	// New in 2023, exact labels to check on ERP and finance extract
	{"BHATNAGAR, AAKARSH", "BHATNAGAR AAKARSH"},
	{"Halagunde,Vinayak Punja", "Halagunde Vinayak Punja"},
}
var cotaPtfOtherPeople = [][]string{
	{"Aguilera Palomino,Diego", "Aguilera Palomino Diego"}, // support from deployment team
	{"Alarcon galvez,Fernando", "Alarcon galvez Fernando"}, // support from deployment team
	{"Bories,Clement", "Bories Clement"},
	{"BHATNAGAR,AAKARSH", "BHATNAGAR AAKARSH"}, // newcomer to replace budgeted resource
	{"Bretagne,Eric", "Bretagne Eric"},         // KMT
	//{"Cabrera,Marcos", "Cabrera Marcos"},
	{"Castano,Esteban", "Castano Esteban"},                 // support from L2
	{"Cerny,Jaroslav", "Cerny Jaroslav"},                   // KMT
	{"Delgado martinez,Alvaro", "Delgado martinez Alvaro"}, // KMT
	{"Fioux,Sebastien", "Fioux Sebastien"},                 // SaaS platform team SM
	{"Freimonas,Romas", "Freimonas Romas"},                 // KMT
	{"Galindo Gomez,Jorge", "Galindo Gomez Jorge"},         // ?
	{"Gukanti,Sandeep", "Gukanti Sandeep"},                 // support from deployment team
	{"Jones,Terence", "Jones Terence"},                     // support from DB expert team
	{"Kumar,Anshuman", "Kumar Anshuman"},                   // support from deployment team
	{"Letolle,Nicolas", "Letolle Nicolas"},                 // support from deployment team
	{"LIM, Mr BERNARD KENNETH", "LIM BERNARD KENNETH"},     // support from DB expert team
	{"Nandiraju,Pavan Kumar", "Nandiraju Pavan Kumar"},     // support from DB expert team
	{"Perez Lagunas,Daniela", "Perez Lagunas Daniela"},     // support from deployment team
	{"Schammass,Alexandre", "Schammass Alexandre"},         // support from deployment team
	{"Sharma,Aditya", "Sharma Aditya"},                     // newcomer to replace budgeted resource
	{"Sharma,Gaurav", "Sharma Gaurav"},                     // deployment/platform team manager
	{"Valette,Karine", "Valette Karine"},                   // support from DB expert team
}

// Delta with baseline/forecast: OTA demo tenant approved by Jérome Voyer [R1R29751]
var ext29751OtaDemoTenantPeople = [][]string{
	{"Virmani,Karan", "Virmani Karan"},
}

// --- R1R29753 Innovation
var innovationBudgetTransPeople = [][]string{
	{"Berard,Xavier", "Berard Xavier"},
	{"Bretagne,Eric", "Bretagne Eric"},
	{"Gattone,Alain", "Gattone Alain"},
	{"Prigent,Francois", "Prigent Francois"},
}
var innovationOtherServerPeople = [][]string{
	{"Castano,Esteban", "Castano Esteban"},
	{"Cerny,Jaroslav", "Cerny Jaroslav"},
	{"Deepak,Deepak", "Deepak Deepak"},
	{"Dhondiyal,Rituraj", "Dhondiyal Rituraj"},
	{"Freimonas,Romas", "Freimonas Romas"},
	{"Kobr,Dan", "Kobr Dan"}, // KMT
	{"Kumar,Anshuman", "Kumar Anshuman"},
	{"Lachowicz,Daniel", "Lachowicz Daniel"},
	{"Miani,Alberto", "Miani Alberto"}, // KMT
	{"Sharma,Gaurav", "Sharma Gaurav"},
	{"Virmani,Karan", "Virmani Karan"},
}
var innovationOtherAppletPeople = [][]string{
	{"Abao,Michael Carlo", "Abao Michael Carlo"},
	{"ABAO,MICHAEL CARLO", "ABAO MICHAEL CARLO"},
	{"Eleserio,Ederlyn", "Eleserio Ederlyn"},
	{"Giva,Joana Marie", "Giva Joana Marie"},
	{"LAI,SER WEI", "LAI SER WEI"},
	{"Marquez,Justin", "Marquez Justin"},
	{"ONG,WILSON LEE", "ONG WILSON LEE"},
}

// --- R1R30027 TAC Dev
var tac2023BudgetServerPeople = [][]string{
	{"Gattone,Alain", "Gattone Alain"},
}
var tacBudgetAppletPeople = [][]string{
	// Applet
	{"Abao,Michael Carlo", "Abao Michael Carlo"},
	{"Eleserio,Ederlyn", "Eleserio Ederlyn"},
	{"ONG,WILSON LEE", "ONG WILSON LEE"},
	{"Shamsudin,Nurrasyidah", "Shamsudin Nurrasyidah"},
}
var tacOtherServerPeople = [][]string{
	// Server
	{"Dumitrescu,Florin", "Dumitrescu Florin"},                                   // support from ODC
	{"Gupta,Ankur", "Gupta Ankur"},                                               // server was not in budget
	{"Hernandez Castaneda,Jose Guillermo", "Hernandez Castaneda Jose Guillermo"}, // support from SSC
	{"Myslivets,Alexey", "Myslivets Alexey"},                                     // support from ODC
	{"Shevnin,Ignat", "Shevnin Ignat"},                                           // support from ODC
	{"Singh,Gurvinder", "Singh Gurvinder"},                                       // server was not in budget
	{"Singhal,Shivank", "Singhal Shivank"},                                       // server was not in budget
}

// Delta with baseline/forecast: Transatel DP+ flow activation for 3k€ [R1R30027]
var ext30027transatelActPeople = [][]string{
	{"Ayasse,Jerome", "Ayasse Jerome"},
}

// --- R1R30028 IOT Dev
var iotBudgetTransPeople = [][]string{
	// Server
	{"ATMOPAWIRO,ALSASIAN", "ATMOPAWIRO ALSASIAN"},
	{"Gattone,Alain", "Gattone Alain"},
}
var iotBudgetAppletPeople = [][]string{
	// Applet
	{"Abao,Michael Carlo", "Abao Michael Carlo"},
	{"Eleserio,Ederlyn", "Eleserio Ederlyn"},
	{"Espinosa,Alen", "Espinosa Alen"},
	{"LAI,SER WEI", "LAI SER WEI"},
}
var iotBudgetServerPeople = [][]string{
	// Server
	{"Kumar,Vishesh", "Kumar Vishesh"},
}
var iotOtherPeople = [][]string{
	// Server
	{"Kumar,Vivek", "Kumar Vivek"},
	{"Prigent,Francois", "Prigent Francois"}, // support from innovation team
	{"Yadav,Sanjeet", "Yadav Sanjeet"},       // support from innovation team
}

var cotaManagersCZ = []string{
	"Franco Mora,Richard Miguel",
	"Fesquet,Sebastien",
	"Kadanik,Jiri",
}
var cotaManagersIN = []string{
	"Pethia,Abhishek",
	"KUMAR,Narendra",
	"Khan,Akram Raza",
	"Kumar,Vikash",
	"Kumar,Vivek",
	"Mehta,Ashish",
	"Gupta,Anshul",
	"Sharma,Gaurav",           // as previous/future manager of COTA resource
	"Gupta,Nainsy",            // as previous/future manager of COTA resource
	"Kanwal,Ishan",            // as previous/future manager of COTA resource
	"Kumar Srivastava,Neeraj", // as previous/future manager of COTA resource
	"Sharma,Ajay",             // as previous/future manager of COTA resource
}
var cotaManagersUS = []string{
	"Letolle,Nicolas",
}
var cotaManagersFR = []string{
	"Pereira Carrari,Mauricio",
	"Vila,Christophe",
}

var peopleExceptions = map[string]string{
	"Barilly,Adrien": "CZ",
}

var cotaManagerCountry pivot.Compute[string] = func(elements []pivot.RawValue) (string, error) {
	if c, ok := peopleExceptions[elements[1].(string)]; ok {
		return c, nil
	}
	for _, m := range cotaManagersCZ {
		if m == elements[0] {
			return "CZ", nil
		}
	}
	for _, m := range cotaManagersFR {
		if m == elements[0] {
			return "FR", nil
		}
	}
	for _, m := range cotaManagersIN {
		if m == elements[0] {
			return "IN", nil
		}
	}
	for _, m := range cotaManagersUS {
		if m == elements[0] {
			return "US", nil
		}
	}
	return fmt.Sprintf("??-%s", elements[0]), nil
}
