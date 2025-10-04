from pydantic import BaseModel,ValidationError


class Person(BaseModel):
    first_name: str 
    last_name:  str
    age:        int
    @property
    def greet_me(self):
        return(f"Welcome, {self.first_name}!! 🤗")


p = Person(first_name='muthu',last_name="kamalan",age=27) # supply as named arguments


print(p) #p.__str__()
print(p.__repr__())  #p.__repr__()
print(f"model fields:: {p.model_fields}")


try:
    p1 = Person(first_name='mariappan')
except ValidationError as ve:
    print(ve)
    # Printing all missing error at same time


print(f"accessing property :: {p.greet_me}")

# since it's dynamic language creating attribute outside of BaseModel
p.age = 10_00_00_00_00_000
print(p)

# changing attribute after it's defined i.e) while defining it's satisfy condition
## remember dynamic language
p.age = '2k24'
print(p)


try:
    # probihited while defining it
    p1 = Person(first_name='mariappan',last_name='m',age='2k24')
except ValidationError as ve:
    print(ve)

# By default, python validate arguments while declarating it. 


