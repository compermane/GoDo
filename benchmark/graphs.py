import matplotlib.pyplot as plt
from collections import Counter
from typing import List


def plot_sequence_count_plot(exp, exp_dur: str, seqs_lens: List[int]) -> None:
        
    count_by_lenght = Counter(seqs_lens)
    lenghts_sorted = sorted(count_by_lenght.items())
    x = [lenght for lenght, count in lenghts_sorted]
    y = [count for lenght, count in lenghts_sorted]

    plt.figure(figsize=(10, 7))
    bars = plt.bar(x, y, color='blue')
    plt.title(f"Distribuição dos tamanhos de sequência ({exp_dur}s)")
    plt.xlabel("Tamanho de Sequência")
    plt.ylabel("Frequência")
    plt.grid(False)
    plt.xticks(x)
    # plt.tight_layout()

    for bar in bars:
        height = bar.get_height()
        plt.text(
            bar.get_x() + bar.get_width() / 2, 
            height + 0.5,                      
            f"{int(height)}",                   
            ha='center', va='bottom', fontsize=9
            )

    plt.savefig(f"{exp.name} - {exp_dur}.png")