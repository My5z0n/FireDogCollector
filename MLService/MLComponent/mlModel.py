from keras.preprocessing.text import Tokenizer
from tensorflow.keras.utils import to_categorical
from tensorflow.keras.models import load_model
from numpy import array
from pickle import dump, load
from keras.preprocessing.text import Tokenizer
from keras.utils import to_categorical
from keras.models import Sequential
from keras.layers import Dense
from keras.layers import LSTM
from keras.layers import Embedding
import numpy as np
from numpy import array

PREDICT_CONST = 0.25


class MLModel:
    modelPath = ""
    tokenizer = None
    model = None

    def __init__(self, modelPath: str = './MLFiredogModel/', ):
        self.modelPath = modelPath

    def loadModel(self, modelName):
        try:
            model = load_model(self.modelPath + modelName+'_model.h5')
            tokenizer = load(
                open(self.modelPath + modelName+'_tokenizer.pkl', 'rb'))

            self.model = model
            self.tokenizer = tokenizer
        except Exception as err:
            return err

    def splitPaths(self, paths):
        lines = []
        length = 2 + 1  # 2 Previous calculate next
        for span in paths:
            for i in range(length, len(span) + 1):
                lines.append(span[i - length:i])
        return lines

    def learn(self, pathsArray, modelName):
        try:
            tokenizer = Tokenizer(oov_token="<OOV>")
            tokenizer.fit_on_texts(pathsArray)
            vocab_size = len(tokenizer.word_index) + 1

            lines = self.splitPaths(pathsArray)

            sequences = tokenizer.texts_to_sequences(lines)
            sequences = array(sequences)

            X, y = sequences[:, :-1], sequences[:, -1]
            y = to_categorical(y, num_classes=vocab_size)
            seq_length = X.shape[1]

            model = Sequential()
            model.add(Embedding(vocab_size, 32, input_length=seq_length))
            model.add(LSTM(32, return_sequences=True))
            model.add(LSTM(32))
            model.add(Dense(16, activation='relu'))
            model.add(Dense(vocab_size, activation='softmax'))
            print(model.summary())

            # compile model
            model.compile(loss='categorical_crossentropy',
                          optimizer='adam', metrics=['accuracy'])

            # fit model
            model.fit(X, y, epochs=32)

            # save the model
            self.model = model
            self.tokenizer = tokenizer
            model.save(self.modelPath + modelName + '_model.h5')
            dump(tokenizer, open(self.modelPath +
                 modelName + '_tokenizer.pkl', 'wb'))
        except Exception as e:
            return e

    def predict(self, paths_array):
        paths = [tmp["span_name"] for tmp in paths_array]
        lines = self.splitPaths(paths)
        sequences = self.tokenizer.texts_to_sequences(lines)
        sequences = array(sequences)

        no_path = 2
        for i in range(len(sequences)):
            no_path += 1
            x = np.asarray([sequences[i][:-1]])
            print(x)
            ret = self.model.predict(x)

            yhat = np.argmax(ret, axis=1)
            # yhat= ret.argsort(axis=1)

            out = sequences[i][-1]
            if ret[0][out] < PREDICT_CONST:
                out_word = "ERR"
                for word, index in self.tokenizer.word_index.items():
                    if index == yhat:
                        out_word = word
                        break
                return True, paths_array[no_path]["span_name"], paths_array[no_path]["span_id"], out_word
                # print(f'{lines[i]} Predicted -> {out_word}')


if __name__ == "__main__":
    paths = [
        ['START', 'A', 'B', 'C', 'D', 'END'],
        ['START', 'A', 'B', 'C', 'F', 'G', 'END'],
        ['START', 'A', 'B', 'H', 'I', 'END'],
    ]

    new = []
    for p in paths:
        for i in range(100):
            new.append(p)
    # print(new)
    x = MLModel()
    x.loadModel()
    x.predict([paths[0]])
