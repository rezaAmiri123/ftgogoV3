Feature: Register Consumer
  Scenario: Registering a new consumer
    Given no consumer named "John Smith" exists
    When I register a new consumer as "John Smith"
    Then I expect the consumr is created
