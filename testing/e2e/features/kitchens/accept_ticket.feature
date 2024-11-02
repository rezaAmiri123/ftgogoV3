Feature: Accept Ticket
  
  Background: 
    Given I am a registered consumer 
    And I sign in
    And I add address to consumer 
    And I create a new restarant
    And I update the restarant menu
    And I create a new order

    
  Scenario: Accept a ticket
    Given no accepted ticket is exists
    When I accept a ticket 
    Then I expect the ticket is accepted

