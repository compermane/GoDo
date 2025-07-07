from typing import List
from const import error_sequences_dirs, non_error_sequences_dirs
from pathlib import Path
import os
import re


padrao = re.compile(r"^Sequence 0:$")

def get_wrong_error_sequences_file(dir: Path) -> List[str]:
    arquivos_errados: List[str] = []

    for nome_arquivo in os.listdir(dir):
        caminho_completo = os.path.join(dir, nome_arquivo)

        if os.path.isfile(caminho_completo):
            try:
                with open(caminho_completo, "r") as f:
                    linhas = f.readlines()

                count_sequence_0 = sum(1 for linha in linhas if "Sequence 0:" in linha)
                if count_sequence_0 > 1 or count_sequence_0 == 0:
                    print(f"Count Sequence 0: {caminho_completo} -> {count_sequence_0}")
                    arquivos_errados.append(nome_arquivo)

            except Exception as e:
                print(f"Erro ao ler {nome_arquivo}: {e}")

    return arquivos_errados

def get_wrong_non_error_sequences_file(dir: Path) -> List[str]:
    linha_alvo = "-------------------Non error sequences-------------------"
    arquivos_invalidos = []

    for nome_arquivo in os.listdir(dir):
        caminho_arquivo = os.path.join(dir, nome_arquivo)

        if not os.path.isfile(caminho_arquivo):
            continue

        try:
            with open(caminho_arquivo, 'r', encoding='utf-8') as f:
                contador = sum(1 for linha in f if linha.strip() == linha_alvo)
                if contador > 1:
                    arquivos_invalidos.append(nome_arquivo)
        except Exception as e:
            print(f"Erro ao ler {nome_arquivo}: {e}")

    return arquivos_invalidos

if __name__ == "__main__":
    with open("wrong.txt", "w") as f:
        for dir in error_sequences_dirs:
            print(f"[error sequences] dir: {dir} -> {get_wrong_error_sequences_file(dir)}", file = f)

        for dir in non_error_sequences_dirs:
            print(f"[non error sequences] dir: {dir} -> {get_wrong_non_error_sequences_file(dir)}", file = f)

    f.close()

