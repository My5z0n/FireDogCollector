from statistics import mode
from keras.preprocessing.text import Tokenizer
from tensorflow.keras.utils import to_categorical
from tensorflow.keras.models import load_model
from numpy import array
from pickle import dump,load
from keras.preprocessing.text import Tokenizer
from keras.utils import to_categorical
from keras.models import Sequential
from keras.layers import Dense
from keras.layers import LSTM
from keras.layers import Embedding

from matplotlib import mlab
import numpy as np
from numpy import array


class MLModel:
    modelPath = ""
    pathsArray = []
    tokenizer = None
    model = None

    def __init__(self, pathsArray,modelPath='./MLFiredogModel/',):
        self.pathsArray=pathsArray
        self.modelPath=modelPath
    

    def loadModel(self):
        model = load_model(self.modelPath+'model.h5')
        tokenizer = load(open(self.modelPath+'tokenizer.pkl', 'rb'))

        self.model = model
        self.tokenizer = tokenizer

    def splitPaths(self, paths):
        lines = []
        length = 2 + 1 #2 Previous calculate next
        for span in paths:
            for i in range(length,len(span)+1):
                lines.append(span[i-length:i])
        return lines

    def learn(self):
        #oov_token="<OOV>"
        tokenizer  = Tokenizer()
        tokenizer.fit_on_texts(self.pathsArray)
        vocab_size = len(tokenizer.word_index) + 1
        #print(t.word_index)

        

        lines = self.splitPaths(self.pathsArray)
        
        sequences  = tokenizer.texts_to_sequences(lines)
        sequences = array(sequences)

        X, y = sequences[:,:-1], sequences[:,-1]
        y = to_categorical(y, num_classes=vocab_size)
        seq_length = X.shape[1]


        model = Sequential()
        model.add(Embedding(vocab_size, 15, input_length=seq_length))
        model.add(LSTM(15, return_sequences=True))
        model.add(LSTM(15))
        model.add(Dense(15, activation='relu'))
        model.add(Dense(vocab_size, activation='softmax'))
        print(model.summary())

        # compile model
        model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])

        # fit model
        model.fit(X, y, epochs=40)

        # save the model 

        self.model = model
        self.tokenizer = tokenizer
        model.save(self.modelPath+'model.h5')
        dump(tokenizer, open(self.modelPath+'tokenizer.pkl', 'wb'))
        model.predi

    def calculateSpan(self,paths):
        lines = self.splitPaths(paths)
        sequences  = self.tokenizer.texts_to_sequences(lines)
        sequences = array(sequences)

        for i in range(len(sequences)):
            x = np.asarray([sequences[i][:-1]])
            print(x)
            ret = self.model.predict(x)
            print(f"Ret: {ret}")
            yhat = np.argmax(ret,axis=1)
            out_word = ''
            for word, index in self.tokenizer.word_index.items():
                if index == yhat:
                    out_word = word
                    break
            print(f'{lines[i]} Predicted -> {out_word}')





if __name__ == "__main__":
    paths = [
        ['START','A','B','C','D','END'],
        ['START','A','B','C','F','G','END'],
        ['START','A','B','H','I','END'],
    ]

    new = []
    for p in paths:
        for i in range(100):
            new.append(p)
    #print(new)
    x = MLModel(new)
    x.loadModel()
    x.calculateSpan([paths[0]])