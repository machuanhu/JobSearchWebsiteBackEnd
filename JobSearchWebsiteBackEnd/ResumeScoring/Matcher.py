import Tire

class Matcher:
    def __init__(self, words, text):
        self.ACautomachine = Tire.Tire(words)
        self.Text = text

    def getResult(self):
        if(len(self.ACautomachine.Search(self.Text))):
            return True
        else:
            return False

    def getNum(self):
        return len(self.ACautomachine.Search(self.Text))