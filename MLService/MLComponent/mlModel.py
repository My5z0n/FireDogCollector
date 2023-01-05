from keras.preprocessing.text import Tokenizer
from tensorflow.keras.utils import to_categorical
from tensorflow.keras.models import load_model
from numpy import array
from pickle import dump, load
from pickle import HIGHEST_PROTOCOL
from keras.preprocessing.text import Tokenizer
from keras.utils import to_categorical
from keras.models import Sequential
from keras.layers import Dense
from keras.layers import LSTM
from keras.layers import Embedding
from keras.layers import Dropout
import numpy as np
from numpy import array
from typing import Tuple

PREDICT_CONST = 0.25


class MLModel:
    modelPath: str = ""
    tokenizer: Tokenizer = None
    model: Sequential = None

    def __init__(self, modelPath: str = './MLFiredogModel/', ):
        self.modelPath = modelPath

    def load_model(self, modelName: str = 'model') -> Exception:
        try:
            model = load_model(self.modelPath + modelName+'_model.h5')
            tokenizer = load(
                open(self.modelPath + modelName+'_tokenizer.pkl', 'rb'))

            self.model = model
            self.tokenizer = tokenizer
        except Exception as err:
            return err

    def split_paths(self, paths: list) -> list:
        lines = []
        length = 2 + 1  # 2 Previous calculate next
        for i in range(length, len(paths) + 1):
            lines.append(paths[i - length:i])
        return lines

    def learn(self, pathsArray, modelName) -> Exception:
        try:
            tokenizer = Tokenizer(oov_token="<OOV>")
            tokenizer.fit_on_texts(pathsArray)
            print(tokenizer.word_index)
            vocab_size = len(tokenizer.word_index) + 1

            lines = self.split_paths(pathsArray)

            sequences = tokenizer.texts_to_sequences(lines)
            sequences = array(sequences)

            X, y = sequences[:, :-1], sequences[:, -1]
            y = to_categorical(y, num_classes=vocab_size)
            seq_length = X.shape[1]

            model = Sequential()
            model.add(Embedding(vocab_size, 32, input_length=seq_length))
            model.add(Dropout(0.2))
            model.add(LSTM(32, return_sequences=True))
            model.add(LSTM(32))
            model.add(Dropout(0.2))
            model.add(Dense(16, activation='relu'))
            model.add(Dense(vocab_size, activation='softmax'))
            print(model.summary())

            # compile model
            model.compile(loss='categorical_crossentropy',
                          optimizer='adam', metrics=['accuracy'])

            # fit model
            model.fit(X, y, epochs=50,batch_size=4)

            # save the model
            self.model = model
            self.tokenizer = tokenizer
            model.save(self.modelPath + modelName + '_model.h5')
            dump(tokenizer, open(self.modelPath +
                 modelName + '_tokenizer.pkl', 'wb'),HIGHEST_PROTOCOL)
        except Exception as e:
            return e

    def predict(self, paths_array) -> Tuple[bool,str,str,str]:
        print(self.tokenizer.word_index)
        paths = [tmp["span_name"].lower() for tmp in paths_array]
        sequences = self.tokenizer.texts_to_sequences([paths])
        sequences = self.split_paths(sequences[0])

        no_path = 2
        for i in range(len(sequences)):
            no_path += 1
            x = np.asarray([sequences[i][:-1]])
            print(x)
            ret = self.model.predict(x)

            yhat = np.argmax(ret, axis=1)

            out = sequences[i][-1]
            if ret[0][out] < PREDICT_CONST:
                out_word = "ERR"
                for word, index in self.tokenizer.word_index.items():
                    if index == yhat:
                        out_word = word
                        break
                return True, paths_array[no_path]["span_name"], paths_array[no_path]["span_id"], out_word

        return False, "", "",""

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
    x.load_model()
    x.predict([paths[0]])
