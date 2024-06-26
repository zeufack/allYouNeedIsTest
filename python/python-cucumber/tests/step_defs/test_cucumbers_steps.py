from pytest_bdd import scenario, given, when,then
from cucumbers.cucumber import CucumberBasket

@scenario('../features/cucumbers.feature', 'Add cucumbers to a basket')
def test_add():
    pass

@given("the basket has 2 cucumbers")
def basket():
    return CucumberBasket(initial_count=2)

@when("4 cucumbers ar added to the basket")
def add_cucumbers(basket):
    basket.add(4)

@then("the basket contains 6 cucumbers")
def step_impl(basket):
    assert basket.count == 6