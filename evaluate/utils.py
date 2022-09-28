import time
from status import TestStatus

class Constraints:
    def __init__(self):
        self.cpu_time = 1.0 # seconds
        self.clock_time = 10.0 # seconds
        self.memmory = 256 # MB

class TestResults:
    def __init__(self, cpu_time:float, clock_time:float, memory_used:float):
        self.cpu_time = cpu_time
        self.clock_time = clock_time
        self.memory_used = memory_used
        
class Test:
    def __init__(self, test_name: str, cons: Constraints):
        self.test_name

class SubmissionResult:
    pass

class Submission:
    def __init__(self, task_name:str, task_vers:str, user_code:str):
        self.timestamp = time.now()
        self.task_name = task_name
        self.task_vers = task_vers
        self.user_code = user_code
    
    def evaluate(self) -> SubmissionResult:
        pass

class Executable:
    pass

def compile(submission: Submission) -> Executable:
    pass