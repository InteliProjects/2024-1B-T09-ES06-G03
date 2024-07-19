import pandas as pd
import numpy as np
from surprise import SVD
from surprise import Dataset
from surprise import Reader
from flask import Flask, request, jsonify
from joblib import dump, load

model_path = "./saveModel/svd_model.joblib"

model = load(model_path)


# Função para transformar os dados da requisição para um DataFrame em Pandas
def transform_to_df(data):
    df = pd.DataFrame(data)
    return df


# Função para retreinar o modelo
def ReTrainModel(df):
    reader = Reader(rating_scale=(-2, 2))
    data = Dataset.load_from_df(df, reader)

    model_ = SVD()
    model_.fit(data)
    dump(model_, model_path)

    global model
    model = model_


# Função para Fazer a predição do modelo
def predictRecomendationModel(df, id_proponente, model=model):
    all_projects = set(df["id_projeto"])  # Todos os projetos
    rated_projects = set(
        df[df["id_proponente"] == id_proponente]["id_projeto"]
    )  # Projetos já avaliados pelo usuário
    projects_to_predict = all_projects - rated_projects  # Projetos não avaliados

    predictions = []
    for project_id in projects_to_predict:
        predictions.append((project_id, model.predict(id_proponente, project_id).est))

    # Ordenar as previsões por estimativa de maior para menor
    predictions.sort(key=lambda x: x[1], reverse=True)
    # Exibir o id dos projetos recomendados
    formatted_array = [
        {"id": item[0]} for item in predictions
    ]
    return formatted_array


app = Flask(__name__)


@app.route("/")
def home():
    return "Hello, Flask!"


@app.route("/train", methods=["POST"])
def train():
    data = request.get_json()
    if not data:
        return jsonify({"error": "No data provided"}), 400

    df = data.get("data")

    validador = df[0]
    name = data.get("name")
    age = data.get("age")

    if (
        not validador.id_proponente
        or not validador.id_projeto
        or not validador.avaliação
    ):
        return (
            jsonify({"error": "id_proponente or id_projeto or avaliação not exist"}),
            400,
        )

    df = transform_to_df(df)
    ReTrainModel(df)

    return jsonify({"message": "Modelo Atualizado"}), 200


@app.route("/predict", methods=["POST"])
def predict():
    data = request.get_json()
    print(data)
    if not data:
        return jsonify({"error": "No data provided"}), 400

    id_proponente_predict = data.get("user_id")
    print(2)

    if not id_proponente_predict:
        return jsonify({"error": "No id_user provided"}), 400

    print(3)
    df = data.get("data")
    print(4)

    validador = df[0]

    if (
        not validador["id_proponente"]
        or not validador["id_projeto"]
        or not validador["avaliacao"]
    ):
        return (
            jsonify({"error": "id_proponente or id_projeto or avaliação not exist"}),
            400,
        )

    df = transform_to_df(df)
    print(4)

    response = predictRecomendationModel(df, id_proponente_predict)
    print(response)

    return jsonify(response), 200


if __name__ == "__main__":
    app.run(debug=True)
