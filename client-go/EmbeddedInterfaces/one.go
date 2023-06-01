package main

import (
	"fmt"
)

type PersonDetailsInt interface {
	PersonDetails()
}

type PersonProgressInt interface {
	PersonProgress()
}

type PersonFinalDetails interface {
	PersonDetailsInt
	PersonProgressInt
}

// type Vikas interface {
// 	PersonDetails()
// 	PersonProgress()
// }

type author struct {
	name            string
	lastName        string
	salary          int
	publish_article int
	total_article   int
}

func (a author) PersonDetails() {
	fmt.Println(a.name)
	fmt.Println(a.lastName)
	fmt.Println(a.publish_article)
	fmt.Println(a.total_article)
}

func (a author) PersonProgress() {
	pending_articles := a.total_article - a.publish_article
	fmt.Println(pending_articles)
}

type scientist struct {
	name                string
	lastName            string
	salary              int
	total_invention     int
	published_invention int
}

func (s scientist) PersonDetails() {
	fmt.Println(s.name)
	fmt.Println(s.lastName)
	fmt.Println(s.published_invention)
	fmt.Println(s.total_invention)
}

func (s scientist) PersonProgress() {
	pending_invention := s.total_invention - s.published_invention
	fmt.Println(pending_invention)
}
