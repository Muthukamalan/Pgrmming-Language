What is Pydantic
- framework for defining specialized python classes.
    - class attributes are called `fields`
    - the specialized class is called `model`
    - serialization and deserialization
    - validate data during deserialization
    - enable us to work in OOP manner


<h6 align='center'>pydantic model=python class</h6>
<div align='center'>
    <img src="./assets/pydantic workflow.png" />
</div>

Pydantic Model VS Data Classes
- both are python classes
- declarative way of defining properties and associated types
- adds easy serialization and deserialization between object and JSON


```py
# Code Generator
from dataclasses import dataclass
@dataclass
class Model:
    lang:str
    version:str
    year:int

# Inheritance property of Class
from pydantic import BaseModel
class Net(BaseModel):
    lang: str
    version: str
    year:int
```