package gamers

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Question struct {
	QuestionText       string
	AnswerOptions      map[int]string
	CorrectAnswerIndex int
}

type Player struct {
	Name  string
	Score int
}

func (p *Player) AddScore() {
	p.Score++
}

func generateQuestions(questionCh chan<- Question) {
	questions := []Question{
		{
			QuestionText: "Чому трюфелі завжди такі дорогі?",
			AnswerOptions: map[int]string{
				1: "Бо вони ростуть під золотими деревами.",
				2: "Бо їх копають мисливці на космічних кораблях.",
				3: "Бо це спосіб трюфельського лобі.",
			},
			CorrectAnswerIndex: 3,
		},
		{
			QuestionText: "Як назвати селянське свято з дегустацією кавунів?",
			AnswerOptions: map[int]string{
				1: "Селянський кавунпразький фестиваль",
				2: "Кавуньянка",
				3: "Водомельонія",
			},
			CorrectAnswerIndex: 1,
		},
		{
			QuestionText: "Що кажуть огірки, коли їх кидають у маринад?",
			AnswerOptions: map[int]string{
				1: "Нам так холодно тут!",
				2: "Це наша часова місія у солонку!",
				3: "Хочемо бути солодшими, не солоними!",
			},
			CorrectAnswerIndex: 2,
		},
		{
			QuestionText: "Які основні причини зміни клімату на планеті?",
			AnswerOptions: map[int]string{
				1: "Це через те, що роботи на місяці змінюють екологію Землі.",
				2: "Відходи від космічних кораблів впливають на атмосферу.",
				3: "Головною причиною є викиди парникових газів людьми.",
			},
			CorrectAnswerIndex: 1,
		},
		{
			QuestionText: "Які історичні факти вплинули на виникнення Ренесансу в Європі?",
			AnswerOptions: map[int]string{
				1: "Ренесанс почався після відкриття кавунпразького фестивалю.",
				2: "Це стало можливим завдяки появі мисливців на космічних кораблях.",
				3: "Період Ренесансу був впливовий через відкриття нових знань і друкованих книг.",
			},
			CorrectAnswerIndex: 3,
		},
		{
			QuestionText: "Які переваги медитації для психічного здоров'я?",
			AnswerOptions: map[int]string{
				1: "Медитація допомагає знайти скарби під золотими деревами.",
				2: "Медитація забезпечує ментальний контакт з інопланетянами.",
				3: "Медитація може знижувати стрес, покращувати концентрацію і зміцнювати психічне здоров'я.",
			},
			CorrectAnswerIndex: 2,
		},
		{
			QuestionText: "Які основні фактори впливають на ціни нафти на світовому ринку?",
			AnswerOptions: map[int]string{
				1: "Ціни нафти залежать від того, які земельні ділянки вивезуть на Місяць.",
				2: "Мисливці на космічних кораблях контролюють ціни на нафту.",
				3: "Попит і пропозиція, політичні конфлікти та виробництво впливають на ціни нафти.",
			},
			CorrectAnswerIndex: 3,
		},
		{
			QuestionText: "Які методи боротьби зі стресом є найефективнішими?",
			AnswerOptions: map[int]string{
				1: "Найкращий спосіб боротьби зі стресом - це відправити огірки на марсіанську відпустку.",
				2: "Медитація і заняття йогою на кораблі під час мисливського сезону.",
				3: "Ефективні методи включають релаксацію, фізичну активність та підтримку від близьких.",
			},
			CorrectAnswerIndex: 1,
		},
		{
			QuestionText: "Які історичні події сприяли розвитку інтернету?",
			AnswerOptions: map[int]string{
				1: "Розвиток інтернету був відповіддю на впровадження технологій водомельонії.",
				2: "Це стало можливим завдяки роботам на Марсі.",
				3: "Інтернет розпочав свій розвиток після винайдення комп'ютерів та створення мережі DARPA.",
			},
			CorrectAnswerIndex: 2,
		},
		{
			QuestionText: "Які фактори впливають на ціни на житло в містах?",
			AnswerOptions: map[int]string{
				1: "Ціни на житло залежать від рівня розвитку космічних технологій.",
				2: "Мисливці на космічних кораблях визначають ціни на нерухомість.",
				3: "Ціни на житло визначаються попитом, пропозицією, регуляційними факторами і економічною ситуацією.",
			},
			CorrectAnswerIndex: 3,
		},
		{
			QuestionText: "Які способи можна застосувати для збереження біорізноманіття?",
			AnswerOptions: map[int]string{
				1: "Для збереження біорізноманіття необхідно посадити дерева під золотими деревами.",
				2: "Мисливці на космічних кораблях відповідають за біорізноманіття на інших планетах.",
				3: "Збереження біорізноманіття включає в себе створення заповідників, регулювання лісорубства та охорону видів.",
			},
			CorrectAnswerIndex: 1,
		},
		{
			QuestionText: "Які технологічні досягнення допомагають у боротьбі зі зміною клімату?",
			AnswerOptions: map[int]string{
				1: "Зміна клімату може бути зупинена завдяки місіям на Місяць.",
				2: "Мисливці на космічних кораблях розвивають технології для збереження клімату.",
				3: "Серед технологічних рішень є використання відновлювальних джерел енергії, зменшення викидів парникових газів та карбонового сліду.",
			},
			CorrectAnswerIndex: 2,
		},
		{
			QuestionText: "Які основні принципи здорового способу життя?",
			AnswerOptions: map[int]string{
				1: "Здоровий спосіб життя включає в себе подорожі на космічних кораблях і вживання кавунів кожного дня.",
				2: "Основними принципами є регулярна фізична активність, збалансоване харчування та відсутність стресу.",
				3: "Здоровий спосіб життя полягає у тому, щоб зберігати огірків у солонці.",
			},
			CorrectAnswerIndex: 2,
		},
	}
	for i := 0; i < len(questions); i++ {
		questionCh <- questions[i]
		time.Sleep(10 * time.Second)
	}

	close(questionCh)
}

func playerAnswering(players *[]Player, questionCh <-chan Question, calculateScoreCh chan<- Player, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("playerAnswering finished!")
			return
		case question, ok := <-questionCh:
			if !ok {
				fmt.Println("questionCh is closed!")
				close(calculateScoreCh)
				return
			}	
			fmt.Println("")
			fmt.Printf("%s (%d)\n", question.QuestionText, question.CorrectAnswerIndex)
			for i, option := range question.AnswerOptions {
				fmt.Printf("%d. %s\n", i, option)
			}
			fmt.Println("")
			for i := range *players {
				randomAnswer := rand.Intn(len(question.AnswerOptions)) + 1
				fmt.Print((*players)[i].Name, ":", randomAnswer, ", ")
				if randomAnswer == question.CorrectAnswerIndex {
					calculateScoreCh <- (*players)[i]
					(*players)[i].AddScore()
				}
			}
			fmt.Println("")
		}
	}
}

func calculate(calculateScoreCh <-chan Player, finishProgramCh chan<- bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("calculate is DONE!")
			finishProgramCh <- true
			return
		case player, ok := <-calculateScoreCh:
			if !ok {
				fmt.Println("calculateScoreCh is closed!")
				finishProgramCh <- true
				return
			}
			player.AddScore()
			fmt.Println(player.Name, "answered correctly!")
		}
	}

}

func Game() {
	fmt.Println("-------=========Game started!=========-------")
	var namesStr string
	flag.StringVar(&namesStr, "names", "Акош Августа Бодоґ Чаба Чолт Ежайяш Ференц Гуґо Жиґмонд Йоужі Бачі", "space-separated list of names")

	flag.Parse()
	names := strings.Split(namesStr, " ")

	players := make([]Player, len(names))

	for i, name := range names {
		players[i] = Player{Name: name}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	questionCh := make(chan Question)
	calculateScoreCh := make(chan Player)
	finishProgramCh := make(chan bool)

	go generateQuestions(questionCh)
	go playerAnswering(&players, questionCh, calculateScoreCh, ctx, cancel)
	go calculate(calculateScoreCh, finishProgramCh, ctx)

	<-finishProgramCh

	fmt.Println("-------=========Game finished! Results:=========-------")
	for _, player := range players {
		fmt.Printf("%s: %d балів \n", player.Name, player.Score)
	}
}
