# golang-test-driven-development

## Introduction

I wanted to get some hands on experience with golang whilst also consolidating my learning on TDD so I followed this article: https://medium.com/@elliotchance/lets-talk-about-tdd-baby-5ce11bd9b53b BUT I tried it with GO instead of PYTHON! 

## To run the tests for yourself

To be able to run the tests in this repo you will need to do the following:

1. clone this repo
1. cd into the this repo
1. run the command: "brew install go"
1. run the command: "go get github.com/stretchr/testify/assert" (this is equivilant to "pip install xyz")
    1. Make sure you .gitconfig file is up to date for this
1. finally to run the tests run: "go test -v fizz_buzz.go fizz_buzz_test.go"
1. Tests will now run.  