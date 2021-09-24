import WrongDetector
import ListTest

if __name__ == '__main__':
    while True:
        with open("Resume.txt", "r", encoding='utf-8') as f:
            text = f.read()  # 读取文件
        with open("Resume.txt", "w", encoding='utf8') as f2:
            f2.truncate()
        detector = WrongDetector.WrongGrammarDetector(text)
        tester = ListTest.ListTester(text)
        Score = 50 - detector.getWrongNum() + tester.getListTest()
        if Score < 0:
            Score = 0
        if Score > 100:
            Score = 100
        with open("score.txt", "w", encoding='utf-8') as F:
            F.write(str(Score))
