package main

import (
	"bufio"
	"encoding/csv"
	"github.com/cvila84/erpdump/internal/erp"
	"github.com/cvila84/erpdump/pkg/table"
	"github.com/cvila84/erpdump/pkg/utils"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	// records[0]=manager
	// records[1]=employee
	// records[6]=hours
	// records[9]=project
	// records[10]=task
	// records[12-17]=times
	erpTimeCards := readCsvFile("./erp2022.csv")
	employeesTimes := &utils.Vector[erp.EmployeeTimes]{ID: func(element erp.EmployeeTimes) string { return element.Name }}
	for i, card := range erpTimeCards {
		if i > 0 {
			manager := strings.TrimSpace(card[0])
			employee := strings.TrimSpace(card[1])
			employeeTimes, ok := employeesTimes.Get(employee)
			if !ok {
				employeeTimes = &erp.EmployeeTimes{Name: employee, ManagerName: manager}
				employeesTimes.Add(employeeTimes)
			}
			month, hours1, hours2, err := erp.MonthlyHours(card)
			if err != nil {
				panic(err)
			}
			if hours1 > 0 || hours2 > 0 {
				employeeTimes.Add(erp.ParseProjectID(card[9]), card[10], month, hours1, hours2)
			}
		}
	}
	// records[0]=employee
	// records[1]=manager
	// records[2]=project
	// records[3]=task
	// records[4-15]=times
	var rawData [][]interface{}
	for _, data := range employeesTimes.GetAll() {
		rawData = append(rawData, data.GetAll()...)
	}
	otaPeople := []string{
		// IOT
		"ATMOPAWIRO,ALSASIAN",
		"Gattone,Alain",
		// OTA L3
		"Agrawal,Somya",
		"Choubisa,Vidhi",
		"Devanjali,Devanjali",
		"Dubey,Parul",
		"GOYAL,JAIDEV",
		"Gupta,Nainsy",
		"Gupta,Swati",
		"KUMAR,Rahul",
		"Kumar Verma,Ayush",
		"Kumar,Sachin",
		"Kushwaha,Amarjit",
		"Martinez Carino,Conrado",
		"Moudgil,Rhea",
		"Patil,Ramesh",
		"Perez Cuellar,Julio Cesar",
		"RAJU,RAMESH",
		"SAHADEVAN,SANU",
		"Shameem,Bilal",
		"Sharma,Naveen",
		"Srivastava,Sahil",
		// VM POC
		"Dokladal,Jakub",
		"Lachowicz,Daniel",
		"Przytarski,Bartlomiej",
		"Sedlacek,Ondrej",
		// SaaS
		"Cerny,Jaroslav",
		"Deepak,Deepak",
		"Kumar,Anshuman",
		"LAMBA,PREETIKA",
		"Levacher,Frederic",
		"LIM, Mr BERNARD KENNETH",
		"Ragot,Emmanuel",
		"Roucheton,Jerome",
		"Singh,Pradeep",
		"Singh,Yash Pal",
		// Strategic
		"Bories,Clement",
		"Cabrera,Marcos",
		"Chiaramello,Daniel",
		"Gukanti,Sandeep",
		"Letolle,Nicolas",
		"Schammass,Alexandre",
		// France
		"Cabagno,Anne",
		"Tessier,Alexandra",
		"Fioux,Sebastien",
		"Coste,Florent",
		// Other
		"Barilly,Adrien",
		"Oremus,Tomas",
		// Prague
		"Delgado martinez,Alvaro",
		"Gatica Peralta,Elia Azucena",
		"Gorokhov,Nikita",
		"Gorysz,Lukasz",
		"Hamid,Juba",
		"Hernandez Castaneda,Jose Guillermo",
		"Hlavacek,Ludek",
		"Jiricek,Libor",
		"Kalita,Victoria",
		"Kostohryz,Jan",
		"Manasek,Radek",
		"Nguyen Gia Can,Cyril",
		"Norko,Veronika",
		"Penha,Bruno",
		"Pomorski,Marcin",
		"Sramek,Vaclav",
		"Uzun,Eraslan",
		"Vondracek,Martin",
		// Noida
		"Agrawal,Deepak",
		"Agrawal,Ritika",
		"Ali,Riyasat",
		"BANSAL,ANKIT",
		"Bhardwaj,Abhishek",
		"Bhatia,Akhil",
		"Bhatnagar,Manas",
		"Dhondiyal,Rituraj",
		"Chauhan,Yashvant Singh",
		"Gupta,Ankur",
		"Gupta,Anshul",
		"Harshadbhai,Patel Kaushikkumar",
		"Jain,Himanshu",
		"Jain,Shubham",
		"Jha,Ashwani Kumar",
		"Kalani,Anukriti",
		"KANSAL,MANIKA",
		"Khan,Akram Raza",
		"Khan,Momin",
		"Khan,Wasim",
		"Kumar,Amrish",
		"Kumar,Chandan",
		"Kumar,Gaurav",
		"KUMAR,Narendra",
		"Kumar,Satyam",
		"kumar,Saurabh",
		"Kumar,Sujit",
		"Kumar,Vikas",
		"Kumar,Vishesh",
		"Kumar,Vivek",
		"Mehta,Ashish",
		"PANDEY,Priyesh",
		"Pandey,Renu",
		"Pethia,Abhishek",
		"Priyanka,Kumari",
		"Rajwaar,Shweta",
		"Rawat,Yoginder",
		"Rupera,Bhumika",
		"Singh,Gurvinder",
		"SINGH,VIJAYLAXMI",
		"SHARMA,ADITI",
		"Sharma,Ashwani Kumar",
		"Sharma,Himanshu",
		"Sharma,Prateek",
		"Sharma,Roshan Lal",
		"Singh,Akshay",
		"Singh,Ashutosh",
		"Singh,Bheem",
		"Singh,Devyani",
		"Singh,Hameer",
		"Singhal,Shivank",
		"Sirvya,Anshul",
		"Vats,Vishant Kumar",
		"VOHRA,Mitali",
	}
	otaProjects := []string{
		"R1R29750",
		"R1R29751",
		"R0S29752",
		"R1R29753",
		"R0R29754",
		"R1R30027",
		"R1R30028",
	}
	functionalProjects := []string{
		"RDX0000A",
		"RDX0000H",
		"RDX0000S",
		"RDX0000T",
		"RDX000PT",
	}
	//otaManagers := []string{
	//	"Pereira Carrari, Mr Mauricio",
	//	"Letolle,Nicolas",
	//	"Vila,Christophe",
	//	"Kumar,Vikash",
	//	"Franco Mora,Richard Miguel",
	//	"Fesquet,Sebastien",
	//	"Pethia,Abhishek",
	//	"Khan,Akram Raza",
	//	"Mehta,Ashish",
	//	"KUMAR,Narendra",
	//	"Gupta,Anshul",
	//	"Kumar,Vivek",
	//}
	tcd := table.NewFloatTable(rawData).
		//Filter(1, table.In(otaManagers)).
		Filter(2, table.In(otaProjects)).
		Row([]int{0}, table.Group1(otaPeople, "OTA", "External"), nil, table.AlphaSort).
		StandardRow(0).
		Column([]int{2}, table.Group2(otaProjects, functionalProjects, "OTA", "Functional", "Other"), nil, table.AlphaSort).
		StandardColumn(2).
		//StandardColumn(3).
		Values([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, utils.YearlyHours, table.Sum, nil)
	err := tcd.Generate()
	if err != nil {
		panic(err)
	}
	file, err := os.Create("erp2022-tcd.csv")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(tcd.ToCSV())
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
