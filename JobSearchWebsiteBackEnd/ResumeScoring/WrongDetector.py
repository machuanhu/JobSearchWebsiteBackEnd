# use pycorrector 进行错别字统计
# Interpreter path:C:\Users\MoonLight\AppData\Local\Programs\Python\Python39\python.exe
import pycorrector

import sys

sys.path.append("..")

class WrongGrammarDetector:
    def __init__(self, text):
        self.CorrectedResume, self.detail = pycorrector.correct(text)
        #print(self.detail)

    def getWrong(self):
        return self.CorrectedResume, self.detail

    def getWrongNum(self):
        return len(self.detail)

