from utils import *

user = await compile(submission)
for test in tests:
    with test:
        output = File(role="answer")
        user(stdin=test.input, stdout=output)
        checker(test.input, output, test.answer)