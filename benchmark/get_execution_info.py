import math
import re
import statistics
import csv
from collections import defaultdict
from pathlib import Path
from typing import List, Tuple
from const import error_sequences_dirs, non_error_sequences_dirs, coverages_dirs
# from graphs import plot_sequence_count_plot

class Sequence:
    def __init__(self, lenght: int):
        # self.funcoes = funcoes
        self.lenght  = lenght

class Run:
    def __init__(self, non_error_sequences: List[Sequence], error_sequences: List[Sequence], file: str, coverage: float):
        self.file                = file
        self.error_sequences     = error_sequences
        self.coverage            = coverage
        self.num_error_sequences = len(error_sequences)
        self.error_avg_lenght    = get_avg_sequence_len(error_sequences)
        self.error_stdd_lenght   = get_stdd_sequence_len(error_sequences)
        self.error_max_lenght    = get_max_sequence_len(error_sequences)
        self.error_min_lenght    = get_min_sequence_len(error_sequences)

        self.non_error_sequences     = non_error_sequences
        self.num_non_error_sequences = len(non_error_sequences)
        self.non_error_avg_lenght    = get_avg_sequence_len(non_error_sequences)
        self.non_error_stdd_lenght   = get_stdd_sequence_len(non_error_sequences)
        self.non_error_max_lenght    = get_max_sequence_len(non_error_sequences)
        self.non_error_min_lenght    = get_min_sequence_len(non_error_sequences)

class Experiment:
    def __init__(self, name: str, runs: List[Run]):
        self.name = name
        self.runs = runs
        self.duration = decide_duration(runs[0].file)
        self.error_avg, self.non_error_avg = get_experiment_avg_sequence_len(runs)
        self.max_error_len, self.max_non_error_len = get_experiment_max_sequence_len(runs)
        self.min_error_len, self.min_non_error_len = get_experiment_min_sequence_len(runs)
        self.avg_stdd_error_len, self.avg_stdd_non_error_len = get_experiment_avg_stdd(runs)

    def dump_to_file(self):
        print(f"Starting CSV for {self.name}")
        with open(f"{self.name}.csv", "w", encoding="utf-8") as f:
            header = ["Arquivo/Experimento", "avg_lenght (error)", "min_lenght (error)", "max_lenght (error)", "stdd_lenght (error)", "qtd (error)",
                      "avg_lenght (non_error)", "min_lenght (non_error)", "max_lenght (non_error)", "stdd_lenght (non_error)", "qtd (non_error)", "coverage (%)"]
            writer = csv.DictWriter(f, fieldnames=header)
            writer.writeheader()

            data = [
                {"Arquivo/Experimento": run.file, "avg_lenght (error)": run.error_avg_lenght, "min_lenght (error)": run.error_min_lenght,
                  "max_lenght (error)": run.error_max_lenght, "stdd_lenght (error)": run.error_stdd_lenght, "qtd (error)": len(run.error_sequences),
                  "avg_lenght (non_error)": run.non_error_avg_lenght, "min_lenght (non_error)": run.non_error_min_lenght, "max_lenght (non_error)": run.non_error_max_lenght,
                  "stdd_lenght (non_error)": run.non_error_stdd_lenght, "qtd (non_error)": len(run.non_error_sequences), "coverage (%)": run.coverage}
                for run in self.runs
            ]

            writer.writerows(data)

            data = {
                    "Arquivo/Experimento": exp.name,
                    "avg_lenght (error)": exp.error_avg,
                    "min_lenght (error)": exp.min_error_len,
                    "max_lenght (error)": exp.max_error_len,
                    "stdd_lenght (error)": exp.avg_stdd_error_len,
                    "qtd (error)": exp.sum_qtd_error_seqs(),
                    "avg_lenght (non_error)": exp.non_error_avg,
                    "min_lenght (non_error)": exp.min_non_error_len,
                    "max_lenght (non_error)": exp.max_non_error_len,
                    "stdd_lenght (non_error)": exp.avg_stdd_non_error_len,
                    "qtd (non_error)": exp.sum_qtd_nerror_seqs(),
                    "coverage (%)": exp.get_avg_coverage()
            }

            writer.writerow(data)

    def sum_qtd_nerror_seqs(self):
        sum: int = 0

        for run in self.runs:
            sum += len(run.non_error_sequences)

        return sum
    
    def sum_qtd_error_seqs(self):
        sum: int = 0

        for run in self.runs:
            sum += len(run.error_sequences)

        return sum
    def get_avg_coverage(self):
        sum: float = 0.0

        for run in self.runs:
            sum += run.coverage

        return sum / len(self.runs)

def decide_duration(word: str) -> str:
    if "15s" in word:
        return "15s"
    elif "30s" in word:
        return "30s"
    elif "1min" in word:
        return "60s"
    elif "5min" in word:
        return "300s"
    return "600s"

def get_experiment_avg_stdd(runs: List[Run]) -> Tuple[float, float]:
    error_accum:     float = 0.0
    non_error_accum: float = 0.0

    for run in runs:
        error_accum += run.error_stdd_lenght
        non_error_accum += run.non_error_stdd_lenght

    return (error_accum / len(runs), non_error_accum / len(runs))

def get_experiment_min_sequence_len(runs: List[Run]) -> Tuple[int, int]:
    min_error = math.inf
    min_non_error = math.inf

    for run in runs:
        if run.error_min_lenght < min_error:
            min_error = run.error_min_lenght
        if run.non_error_min_lenght < min_non_error:
            min_non_error = run.non_error_min_lenght

    return (min_error, min_non_error)

def get_experiment_max_sequence_len(runs: List[Run]) -> Tuple[int, int]:
    max_error = -1
    max_non_error = -1

    for run in runs:
        if run.error_max_lenght > max_error:
            max_error = run.error_max_lenght
        if run.non_error_max_lenght > max_non_error:
            max_non_error = run.non_error_max_lenght

    return (max_error, max_non_error)

def get_experiment_avg_sequence_len(runs: List[Run]) -> Tuple[float, float]:
    error_accum:     float = 0.0
    non_error_accum: float = 0.0

    # Calcular a média de error e non_error
    for run in runs:
        error_accum     += run.error_avg_lenght
        non_error_accum += run.non_error_avg_lenght

    return (error_accum / len(runs), non_error_accum / len(runs))

def get_stdd_sequence_len(sequences: List[Sequence]) -> float:
    lenghts: List[int] = []

    for seq in sequences:
        lenghts.append(seq.lenght)

    return statistics.stdev(lenghts)

def get_min_sequence_len(sequences: List[Sequence]) -> int:
    min = math.inf

    for seq in sequences:
        if seq.lenght < min:
            min = seq.lenght

    return int(min)

def get_max_sequence_len(sequences: List[Sequence]) -> int:
    max = 0

    for seq in sequences:
        if seq.lenght > max:
            max = seq.lenght

    return max

def get_avg_sequence_len(sequences: List[Sequence]) -> float:
    sum: int = 0

    for seq in sequences:
        sum += seq.lenght

    return sum / len(sequences)

def get_all_files_from_dir(dir: Path) -> List[str]:
    directory = Path(dir)

    files = [f for f in directory.iterdir() if f.is_file()]
    return files

def get_all_sequences_from_file(file: str) -> List[Sequence]:
    seqs: List[Sequence] = []

    with open(file, "r") as f:
        for line in f:
            funcoes: List[str] = re.findall(r'\b(\w+)\s*\(', line)
            if len(funcoes) > 0:
                seqs.append(Sequence(len(funcoes)))

    return seqs

def extract_repo_and_time_from_dir(dir_path: str) -> tuple[str, str]:
    parts = Path(dir_path).parts
    repo = parts[-4]  
    time = parts[-2].replace("_runs", "")  
    return repo, time

def extract_repo_and_time_from_coverage_dir(dir_path: Path) -> tuple[str, str]:
    parts = Path(dir_path).parts
    repo = parts[-4]  # cobra
    time = parts[-2].replace("_run_info", "")  # 15s, 1min etc
    return repo, time

def get_all_coverages_from_dir(info_dir: Path) -> List[float]:
    files = sorted(Path(info_dir).glob("*"))
    coverages = []

    for file in files:
        coverage = extract_coverage_from_file(file)
        coverages.append(coverage)

    return coverages

def extract_coverage_from_file(file_path: Path) -> float:
    with open(file_path, "r") as f:
        for line in f:
            match = re.search(r"coverage:\s+([\d.]+)%", line)
            if match:
                # print(f"{file_path}: {float(match.group(1))}")
                return float(match.group(1))
            
    return -1.0

if __name__ == "__main__":
    experiments = []

    grouped_dirs  = defaultdict(lambda: {"non_error_dir": [], "error_dir": []})
    coverages_map = {}

    # Coleta dos non-error
    for dir_path in non_error_sequences_dirs:
        print(f"{dir_path}: {type(dir_path)}")
        repo, time = extract_repo_and_time_from_dir(dir_path)
        grouped_dirs[(repo, time)]["non_error_dir"] = Path(dir_path)

    # Preencher error dirs
    for dir_path in error_sequences_dirs:
        print(f"{dir_path}: {type(dir_path)}")
        repo, time = extract_repo_and_time_from_dir(dir_path)
        grouped_dirs[(repo, time)]["error_dir"] = Path(dir_path)

    # Obter os coverages:
    for dir_path in coverages_dirs:
        print(f"{dir_path}: {type(dir_path)}")
        repo, time = extract_repo_and_time_from_coverage_dir(dir_path)
        coverages = get_all_coverages_from_dir(dir_path)
        coverages_map[(repo, time)] = coverages

    # Montar experiments
    for (repo, time), dirs in grouped_dirs.items():
        non_error_dir = dirs["non_error_dir"]
        error_dir = dirs["error_dir"]

        if not non_error_dir or not error_dir:
            print(f"[!] Diretórios ausentes para {repo}-{time} {non_error_dir} {error_dir}")
            continue

        # Obter os 30 arquivos de cada pasta
        non_error_files = sorted(get_all_files_from_dir(non_error_dir))
        error_files = sorted(get_all_files_from_dir(error_dir))

        if len(non_error_files) != 30 or len(error_files) != 30:
            print(f"[!] Esperado 30 execuções para {repo}-{time}, mas recebeu {len(non_error_files)} non-error e {len(error_files)} error")
            continue
        
        print(coverages_map)
        coverages = coverages_map.get((repo, time))
        if not coverages or len(coverages) != 30:
            print(f"[!] Sem coverages ou número incorreto para {repo}-{time}")
            continue

        runs = []

        for i in range(30):
            non_error_seqs = get_all_sequences_from_file(non_error_files[i])
            error_seqs = get_all_sequences_from_file(error_files[i])
            cov = coverages[i]
            run_label = f"{repo}/{time}/run{i}"
            runs.append(Run(non_error_seqs, error_seqs, run_label, cov))

        experiments.append(Experiment(f"{repo}-{time}", runs))
        
    for exp in experiments:
        exp.dump_to_file()
