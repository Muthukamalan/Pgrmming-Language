# Deserialization:: Load data into Pydantic Model

from pydantic import BaseModel,ValidationError
import json


data:dict = {
    'first_name':'Richard',
    'last_name':'Feynman',
    'age':64
}


class Person(BaseModel):
    first_name: str 
    last_name:  str 
    age:        int
    @property
    def greet_me(self):
        return(f"Welcome, {self.first_name}!! 🤗")



 
# Pydantantic model loads from JSON/dict
'''
with open('person.json', 'w') as json_file: 
    json_file.write(json.dumps(json_data))

# convert JSON to dict
with open('person.json','r') as f:
    json.loads(f)
'''



p = Person(first_name='muthu',last_name="kamalan",age=27) # supply as named arguments
p0 = Person(**data)   # inefficient when nested json occurs


p1 = Person.model_validate(data)  # << passing as dict
print(p1)

try:
    Person.model_validate({ 'first_name':'Richard',    'age':64 })
except ValidationError as ve:
    print(ve.__class__)



print(json.dumps(data),type(json.dumps(data)))  # convert dict to str

p2 = Person.model_validate_json(json.dumps(data))   # << passing as json-string
print(p2)