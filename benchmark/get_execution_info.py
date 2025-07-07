import math
import re
import statistics
import time
from pathlib import Path
from typing import List, Tuple
from const import error_sequences_dirs, non_error_sequences_dirs

class Sequences:
    def __init__(self, sequences_lens: List[int], id: str):
        self.sequences_lens = sequences_lens
        self.define_average_len()
        self.define_stdd_len()
        self.define_max_len()
        self.define_min_len()
        self.id = id

    def string(self):
        return f"Average: {self.average_len}\nStandard Deviation: {self.stdd_len}\nMax Lenght: {self.max_len}\nMin Lenght: {self.min_len}"

    def define_average_len(self):
        qtd: int = len(self.sequences_lens)
        accum: int = 0

        for seq in self.sequences_lens:
            accum += seq

        if qtd == 0:
            self.average_len = 0
        else:
            self.average_len = accum / qtd

    def define_stdd_len(self):
        self.stdd_len: float = statistics.stdev(self.sequences_lens)

    def define_max_len(self):
        max: int = 0

        for seq in self.sequences_lens:
            if seq > max:
                max = seq

        self.max_len = max

    def define_min_len(self):
        min: int = math.inf

        for seq in self.sequences_lens:
            if seq < min and seq > 2:
                min = seq

        self.min_len = min

    def dump_to_file(self, mode: str):
        if mode == "error":
            file_name = self.id + "-error-sequences-statistics.txt"
        else:
            file_name = self.id + "-non-error-sequences-statistics.txt"

        with open(file_name, "w") as f:
            print(f"Average Sequence Len: {self.average_len}", file = f)
            print(f"Standard Deviation Sequence Len: {self.stdd_len}", file = f)
            print(f"Max Sequence Len: {self.max_len}", file = f)
            print(f"Min Sequence Len: {self.min_len}", file = f)

def get_all_coverages(dir: str) -> None:
    pass

def get_all_sequences(file: Path) -> Sequences:
    seq_lens: List[int] = []

def get_runs_information(dir: str) -> Tuple[Sequences, str]:
    directory = Path(dir)
    files = [f for f in directory.iterdir() if f.is_file()]
    seq_lens: List[int] = []
    id: str = "undefined-id"

    for file in files:
        parts: Tuple[str, ...] = file.parts
        id                     = parts[2] + "-" + parts[4]
        with open(file, "r") as f:
            for line in f:
                funcoes: List[str] = re.findall(r'\b(\w+)\s*\(', line)
                if len(funcoes) > 0:
                    seq_lens.append(len(funcoes))

            f.close()

    return Sequences(seq_lens, id)

if __name__ == "__main__":
    start = time.time()
    err_seqs: List[Sequences] = []
    non_err_seqs: List[Sequences] = []

    for dir in error_sequences_dirs:
        err_seqs.append(get_runs_information(dir))

    for dir in non_error_sequences_dirs:
        non_err_seqs.append(get_runs_information(dir))

    for seq in err_seqs:
        seq.dump_to_file("error")

    for seq in non_err_seqs:
        seq.dump_to_file("non_error")

    end = time.time()
    print(f"Produced statistics for {len(err_seqs) + len(err_seqs)} executions in {end - start}")