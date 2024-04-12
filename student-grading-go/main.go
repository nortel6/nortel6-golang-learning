package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// I am just gonna assume there will be no need
// to further validate here
// Good enough
func parseInt(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}

func parseCSV(filePath string) []student {
	// 30 students in the file but I don't want
	// to make it into an assumption here
	var students []student

	/*
	 * I looked around, and it seems like the choice to
	 * read file is either sequential or all at once.

	 * In this case, I can probably go for io.Open(),
	 * then split the bytes retrieved with a \n.

	 * But, I will just assume the file can be big, so
	 * I am just gonna use bufio Scanner
	 */
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// So that I don't forget later
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Skipping the headers row
	scanner.Scan()

	// Scanner is sufficient here, unless
	// there exists a line with over 64k character
	for scanner.Scan() {
		// I can handle it as bytes instead, but nvm
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		// The marks have to be parsed separately
		// Note: I am not gonna add extra validation to this

		// Just realized you can choose which
		// field to assign the values to
		// Good if you don't want to be coupled
		// on the order/removal/addition of fields
		students = append(students, student{
			firstName:  splitLine[0],
			lastName:   splitLine[1],
			university: splitLine[2],
			test1Score: parseInt(splitLine[3]),
			test2Score: parseInt(splitLine[4]),
			test3Score: parseInt(splitLine[5]),
			test4Score: parseInt(splitLine[6]),
		})
	}

	// Scanner could stop because of error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return students
}

func gradeScore(score float32) Grade {
	switch {
	case score < 35:
		return F
	case score >= 35 && score < 50:
		return C
	case score >= 50 && score < 70:
		return B
	// case score >= 70:
	// Must be default or else it complains
	default:
		return A
	}
}

func calculateGrade(students []student) []studentStat {
	// I will just allocate at once instead of appending
	studentStats := make([]studentStat, len(students))

	for i, student := range students {
		score := float32(student.test1Score+
			student.test2Score+
			student.test3Score+
			student.test4Score) / 4
		studentStats[i] = studentStat{
			student:    student,
			finalScore: score,
			grade:      gradeScore(score),
		}
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	top := studentStat{finalScore: 0}
	for _, s := range gradedStudents {
		if s.finalScore > top.finalScore {
			top = s
		}
	}
	return top
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	tops := make(map[string]studentStat)
	for _, s := range gs {
		u := s.university
		// If key does not exist or the final score is better than the previous student
		if top, ok := tops[u]; !ok || s.finalScore > top.finalScore {
			tops[u] = s
		}
	}
	return tops
}
