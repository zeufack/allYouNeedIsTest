Feature: Cucumber Basket

    As a gardener,
    I want to carry cucumbers in as basket,
    So that I don't drop them all.

    Scenario: Add cucumbers to a basket
        Given the basket has 2 cucumbers
        When 4 cucumbers ar added to the basket
        Then the basket contains 6 cucumbers