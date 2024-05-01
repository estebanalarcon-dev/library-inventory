Feature: borrow a book
  In order to do a borrow
  As a user
  I need to borrow a book
  Scenario: Borrowing book
    Given user and book
    When request the borrowing
    Then book is borrowed to user