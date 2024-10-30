Feature: Create order
  
  As a consumer I can create a new order 

  Background: 
    Given I am a registered consumer 
    And I sign in
    And I add address to consumer 
    And I create a new restarant
    And I update the restarant menu

    
  Scenario: Create an order
    Given no order for registered consumer exists
    When I create a new order 
    Then I expect the order is created

