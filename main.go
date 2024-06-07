package main

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Education struct {
	University string
	Department string
}

type Workplace struct {
	Company     string
	Duration    string
	Position    string
	Description string
}

type Certificate struct {
	Name  string
	Image string
	Link  string
}

type WorkStack struct {
	Languages  []string
	Frameworks []string
	Databases  []string
	Others     []string
}

type Project struct {
	Name        string
	Description string
	Screenshots []string
	Readme      string
	GitHubLink  string
}
type AboutMe struct {
	Name           string
	Specialization string
	Location       string
	Age            string
}
type Contact struct {
	GitHub   string
	LinkedIn string
	Email    string
	Telegram string
}

var projects = []Project{
	{
		Name: "Infrastructure monitoring service ",
		Description: `Creation of a microservice for monitoring the status of servers, virtual machines and stands, integration with a telegram
		bot for data output`,
		Screenshots: []string{"/MonitoringService/workMessage1.png", "/MonitoringService/workMessage2.png", "/MonitoringService/workMessage3.png", "/MonitoringService/configFile.png"},
		Readme: `This project was written during an internship at Links Technologies. Its purpose is to check the stands, servers and virtual machines for their condition and workload. 
		The data is displayed in the telegram bot. A customizable system of reminders for inspections has also been implemented.It is also possible to upload a config file.`,
		GitHubLink: "https://github.com/Feof1l/infrastructure-tracking-service/tree/ver4.0",
	},
	{
		Name: "LinkBox",
		Description: `Developing a website for storing notes using a template for the front. The notes were stored in a database
		MySQL with automatic deletion after the expiration date.`,
		Screenshots: []string{"/LinkBox/LinkBoxUi.jpeg", "/LinkBox/LinkBoxDB.png"},
		Readme: `This site is an educational project. Templates were used to write the front.
		A custom logger has been added, a caching system, and you can run the application on different ports using flags from the terminal.
		The server part is written in pure golang. When writing, 
		I tried to take into account all the principles of a clean file organization.The notes are stored in a mysql database. The connection is made using a connection pool. Each note has a lifetime, 
		after which the note is deleted from the table in the database`,
		GitHubLink: "https://github.com/Feof1l/LinkBox",
	},
	{
		Name:        "HrBot",
		Description: `Telegram hr bot that conducts an interview with a candidate for a job instead of hr`,
		Screenshots: []string{"/HrBot/BotInfo.png", "/HrBot/startOfDialog.png", "/HrBot/messageFromBot1.png", "/HrBot/messageFromBot2.png", "/HrBot/Db.png"},
		Readme: `This project is my final qualifying work.The main goal of this bot is to optimize the work of HR.With the help of a bot, the communication time of an employee (a living person) with a candidate for a vacancy will be reduced.Instead, 
		the candidate will conduct a dialogue-correspondence via the social network messenger (telegram) with the bot.Also, the necessary data will be automatically entered into the database. 
		This bot will allow you to screen out candidates who do not meet the requirements of the company, which will reduce the working time and effort of the employee conducting the interview. This employee will conduct a conversation and check the knowledge of candidates who have corresponded with the bot, that is, with those whom the bot has not eliminated.
		Thus, this candidate meets the minimum requirements of the hiring company.`,
		GitHubLink: "https://github.com/Feof1l/TelegramHrBot",
	},
	{
		Name:        "WeatherApp",
		Description: `A simple website for viewing the weather. To get information about the weather, you need to enter the desired city. The application works with the open weather api external API.`,
		Screenshots: []string{"/WeatherApp/startWindow.png", "/WeatherApp/workWindow1.png", "/WeatherApp/workWindow2.png", "/WeatherApp/errorWindow.png"},
		Readme:      `A simple website for viewing the weather. To get information about the weather, you need to enter the desired city. The application works with the open weather api external API.`,
		GitHubLink:  "https://github.com/Feof1l/WeatherApp",
	},
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "My Portfolio",
			"AboutMe": AboutMe{
				Name:           "Vasilov Ivan",
				Specialization: "Go-backend Developer",
				Location:       "Moscow, Russia",
				Age:            strconv.Itoa(time.Now().Year() - 2003),
			},
			"Contacts": Contact{
				GitHub:   "https://github.com/Feof1l",
				LinkedIn: "https://www.linkedin.com/in/ivan-vasilov-85a90a2bb/",
				Email:    "Ivan.vasilov.03@mail.ru",
				Telegram: "https://t.me/feof1l",
			},
			/*"Name":           "Vasilov Ivan",
			"Specialization": "Go-backend Developer",
			"Location":       "Moscow, Russia",*/
			"Education": Education{
				University: "Moscow Power Engineering Institute ( MPEI )",
				Department: "Applied Mathematics and AI",
			},
			"Workplaces": []Workplace{
				{
					Company:     "Links Technologies",
					Duration:    "Jan 2023 - Aug 2023",
					Position:    "Intern backend-developer",
					Description: "Development and support of backend systems.",
				},
				{
					Company:     "Technologies XIX",
					Duration:    "Apr 2024 - Present",
					Position:    "Backend-developer",
					Description: "Improvement and support of existing systems and microservices.",
				},
			},
			"Certificates": []Certificate{
				{
					Name:  "Basics of computer networks",
					Image: "Сертфиикат Coursera Комп. сети.jpg",
					Link:  "https://drive.google.com/file/d/1dW-DiQxROCj6icy0wTVmTJ68AwuBm90A/view",
				},
				{
					Name:  "Programming in golang",
					Image: "Сертификат степик Го.jpg",
					Link:  "https://stepik.org/cert/2027455",
				},
				{
					Name:  "Algorithms and data structures",
					Image: "Сертификат ВК_Алгоритмы.jpg",
					Link:  "https://education.vk.company/curriculum/certificates/download/43156/d9c93934-b2bd-4742-94b9-4a3cde6e847c/",
				},
				{
					Name:  "Yandex CodeBattle",
					Image: "Сертфиикат Яндекс_КодБаттл.jpg",
					Link:  "https://drive.google.com/drive/folders/1PAhw6D6u3Ql1g3a25aoYfQB2RdouMgS5",
				},
			},
			"WorkStack": WorkStack{
				Languages:  []string{"Go", "Python"},
				Frameworks: []string{"Gin", "React", "Django"},
				Databases:  []string{"MySQL", "PostgreSQL"},
				Others:     []string{"Docker", "Git"},
			},
			"Projects": projects,
			"Languages": []map[string]string{
				{"Language": "Russian", "Proficiency": "Native"},
				{"Language": "English", "Proficiency": "Intermediate ( B1 )"},
			},
		})
	})

	router.GET("/project/:name", func(c *gin.Context) {
		name := c.Param("name")
		var project Project
		for _, p := range projects {
			if p.Name == name {
				project = p
				break
			}
		}
		c.HTML(http.StatusOK, "project.html", gin.H{
			"Title":   project.Name,
			"Project": project,
		})
	})

	router.Run(":9090")
}

/*
<section id="projects">
            <h2>Projects</h2>
            <div class="projects-container">
                {{ range .Projects }}
                <div class="project">
                    <h3>{{ .Name }}</h3>
                    <p>{{ .Description }}</p>
                    <a href="/project/{{ .Name }}" class="button">View Project</a>
                </div>
                {{ end }}
            </div>
        </section>

*/
