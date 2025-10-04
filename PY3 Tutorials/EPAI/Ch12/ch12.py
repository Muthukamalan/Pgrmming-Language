from typing import List,Sequence,NamedTuple,Dict
import os 
from datetime import datetime
from collections import namedtuple
import pickle 
from rich import print


cars:List[NamedTuple] = []
# lazy read
with open(os.path.join(os.getcwd(),"assets",'nyc_parking_tickets.csv'),mode='r') as f:
    # call once and pick headers
    header = [_.strip().replace(' ','_') for _ in next(f).split(',')]        
    # Tuple
    Cars = namedtuple(typename='Cars',field_names=header)

    # Read line by line
    for row in f.readlines():
        # print(row)

        if len(row.split(','))==9:
            c  = Cars(*[i for i in row.split(',')])
            c  = c._replace(
                        Issue_Date = datetime.strptime(c.Issue_Date, '%m/%d/%Y'),
                        Summons_Number=int(c.Summons_Number),
                        Violation_Code=int(c.Violation_Code),
                        Violation_Description=c.Violation_Description.strip()
                )
            cars.append(c)
        else:
            ...



with open(os.path.join(os.path.dirname(__file__),'assets','cars.pkl'),mode='wb') as f:
    pickle.dump(cars,f)



# cars
with open(os.path.join(os.path.dirname(__file__),'assets','cars.pkl'),mode='rb') as f:
    cars = pickle.load(f)


# .value_counts
car_make:Dict = {}
for c in cars:
    if c.Vehicle_Make in car_make:
        car_make[c.Vehicle_Make] += 1
    else:
        car_make[c.Vehicle_Make]=0
    
print(car_make)