package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Team struct {
	name    string
	hacks   int
	penalty int
}

// Структура для хранения записи о запросе команды
// type Record struct {
// 	teamName  string
// 	timestamp time.Time
// 	server    string
// 	result    string
// }

func calcPenalty(start time.Time, eventTime time.Time) int {
	return int(eventTime.Sub(start).Minutes())
}

func sortTeams(tmap map[string]*Team) []Team {
	var teams []Team
	for _, team := range tmap {
		teams = append(teams, *team)
	}
	// Сортировка команд по количеству взломанных серверов, затем по штрафному времени, затем по имени
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].hacks != teams[j].hacks {
			return teams[i].hacks > teams[j].hacks
		}
		if teams[i].penalty != teams[j].penalty {
			return teams[i].penalty < teams[j].penalty
		}
		return teams[i].name < teams[j].name
	})

	return teams
}

func Fifth() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение времени начала хакатона
	hackathonTimeStr, _ := reader.ReadString('\n')
	hackathonTimeStr = strings.TrimSpace(hackathonTimeStr)
	hackathonTime, _ := time.Parse("15:04:05", hackathonTimeStr)

	// Чтение количества запросов
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// Хранение результатов для каждой команды
	teamResults := make(map[string]*Team)

	// Чтение запросов
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Split(line, " ")

		teamName := parts[0]
		teamName = strings.Trim(teamName, "\"")
		eventTimeStr := parts[1]
		eventTime, _ := time.Parse("15:04:05", eventTimeStr)
		result := parts[3]
		result = strings.TrimSpace(result)

		if result == "ACCESSED" {
			penalty := calcPenalty(hackathonTime, eventTime)

			// Добавляем отладочный вывод перед изменением карты

			if _, exists := teamResults[teamName]; !exists {
				teamResults[teamName] = &Team{name: teamName}
			}

			teamResults[teamName].hacks++
			teamResults[teamName].penalty += penalty

		} else if result == "FORBIDEN" || result == "DENIED" {
			if _, exists := teamResults[teamName]; !exists {
				teamResults[teamName] = &Team{name: teamName}
			}

			// Исправляем: приводим штраф к секундам
			teamResults[teamName].penalty += 20 // 20 минут = 1200 секунд
		}
	}
	teams := sortTeams(teamResults)
	place := 1
	for i := range teams {
		if i > 0 && (teams[i].hacks != teams[i-1].hacks || teams[i].penalty != teams[i-1].penalty) {
			place = i + 1
		}
		fmt.Printf("%d \"%s\" %d %d\n", place, teams[i].name, teams[i].hacks, teams[i].penalty)
	}
}
