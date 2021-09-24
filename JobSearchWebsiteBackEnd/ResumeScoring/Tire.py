from collections import defaultdict

class TireNode(object):
    def __init__(self, value=None):
        self.value = value
        self.fail = None
        self.tail = 0
        self.children = {}

class Tire(object):
    def __init__(self, words):
        self.root = TireNode()
        self.count = 0
        self.words = words
        for word in words:
            self.Insert(word)
        self.FindFail()

    def Insert(self, sequence):
        self.count = self.count + 1
        currentNode = self.root
        for item in sequence:
            if item in currentNode.children:
                currentNode = currentNode.children[item]
            else:
                child = TireNode(value=item)
                currentNode.children[item] = child
                currentNode = child
        currentNode.tail = self.count

    def FindFail(self):
        queue = [self.root]
        while len(queue):
            currentNode = queue[0]
            del queue[0]
            for value in currentNode.children.values():
                if currentNode == self.root:
                    value.fail = self.root
                else:
                    p = currentNode.fail
                    while p:
                        if value.value in p.children:
                            value.fail = p.children[value.value]
                            break
                        p = p.fail
                    if not p:
                        value.fail = self.root
                queue.append(value)

    def Search(self, text):
        p = self.root
        start_index = 0
        FindList = defaultdict(list)
        for i in range(len(text)):
            single_char = text[i]
            # print(single_char)
            while single_char not in p.children and p is not self.root:
                p = p.fail
            if single_char in p.children and p is self.root:
                start_index = i
            if single_char in p.children:
                p = p.children[single_char]
            else:
                start_index = i
                p = self.root
            temp = p
            while temp is not self.root:
                if temp.tail:
                    FindList[self.words[temp.tail - 1]].append((start_index, i))
                temp = temp.fail
        return FindList