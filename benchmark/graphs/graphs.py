import matplotlib.pyplot as plt
from matplotlib.ticker import MaxNLocator

experiment_times = [
    15, 30, 60, 300, 600
]

repo_names = [
    "cobra",
    "dive",
    "echo",
    "esbuild",
    "fiber",
    "fzf",
    "gin",
    "gorm",
    "harness",
    "kit",
    "logrus",
    "nsq"
]

avg_lenght_err_stats = [
    # cobra
    [6.154380584138095, 6.4670772566164665, 7.26166528451624, 8.335095809561592, 8.697788610571306],

    # echo
    [6.84734173437737, 7.376868970634918, 7.535272593720066, 8.550997250961688, 9.386680090648975],

    # dive
    [6.583896033152028, 6.9256957463955535, 7.642460650227659, 8.751511815335222, 9.338256167957772],

    # esbuild
    [4.841696801703795, 5.369939645377031, 6.176304724950394, 7.613845767307241, 8.278487337788498],

    # fiber
    [5.175766050077873, 6.235013604892933, 6.363402353649286, 8.049661558374728, 8.507713479882053],

    # fzf
    [5.167652483435963, 5.6641874124140035, 6.294472641330391, 7.990035282130462, 8.460623917161746],

    # gin
    [3.4761742146481565, 3.9821166205626977, 4.836863386360618, 6.4421741415574925, 7.023028337075365],

    # gorm
    [5.463452771902054, 6.146156371580986, 6.621044294066186, 7.607593175861277, 7.940533067956117],

    # harness
    [5.47679489773468, 5.8069460490963065, 6.492892154718025, 7.680468349499004, 8.072539159682425],

    # kit
    [6.583606570361831, 7.25631618759509, 7.875537398134695, 8.83582608825884, 9.254551769659468],

    # logrus
    [7.6206589165002585, 8.014228525493822, 8.312017288375069, 9.565799716261308, 9.812056180629591],

    # nsq
    [4.628639937923123, 4.892588088449075, 5.599157035992929, 6.859610033143336, 7.092488251227097]

]

avg_lenght_nerr_stats = [
    #cobra
    [6.15916464319774, 6.4670772566164665, 7.26166528451624, 8.335095809561592, 8.697788610571306],

    # echo
    [6.888605058201531, 7.450650989736908, 7.584149961579432, 8.605136283959611, 9.456834244665522],

    # dive
    [6.664263008039634, 6.948813887708785, 7.7238179132840985, 8.845857750051886, 9.465692492692103],

    # esbuild
    [4.821125331091198, 5.390329926698542, 6.19456954949496, 7.612362911024844, 8.294125225034618],

    # fiber
    [5.383216817262407, 6.470211814533548, 6.575512484515799, 8.268674002135409, 8.691945642204923],

    # fzf
    [5.187183532257564, 5.731744768601569, 6.310835462390596, 8.05190782417844, 8.508274482406565],

    # gin
    [3.3786104396421193, 3.940015033117872, 4.804277426723702, 6.427725742708192, 7.019378106360144],

    # gorm
    [5.420742817349891, 6.108279414527031, 6.555451261407272, 7.57284130402606, 7.925614597029363],

    # harness
    [5.514203011555627, 5.7948680063809785, 6.489805771358879, 7.694373166853845, 8.087740562440418],

    # kit
    [6.588454146893608, 7.287836170168665, 7.8851363910330265, 8.855863687070434, 9.264111813500007],

    # logrus
    [7.796690553017757, 8.18657036493744, 8.480640038523793, 9.711568776963137, 9.967136078116702],

    # nsq
    [4.710976922519523, 4.958412434702471, 5.640234875381631, 6.842042498057717, 7.058149313443154]
]

avg_stdd_err_stats = [
    # cobra
    [2.2249214929323506, 2.3215451842997465, 2.547057711860256, 2.6804258192368082, 2.7117242688529113],

    # echo
    [2.513626563405572, 2.5947307717536896, 2.6017832969341725, 2.733962564345606, 2.9704843321204684],

    # dive
    [2.439760055818617, 2.4308924551668403, 2.653028431280269, 2.889288471081859, 3.022166445995325],

    # esbuild
    [1.9192148383827887, 1.9828178790201432, 2.300404580336986, 2.553976969189939, 2.687143490151904],

    # fiber
    [2.208014856277649, 2.6120095517228417, 2.71424891456805, 3.1352415135351897, 3.285624508361993],

    # fzf
    [2.064119403251569, 2.200578440457433, 2.275332614509745, 2.65190666871755, 2.747055207093363],

    # gin
    [1.4969724792989136, 1.7059278109114566, 1.9413307867303493, 2.3383759895667633, 2.394264564281879],

    # gorm
    [2.0212836734672965, 2.211393933803696, 2.290371856407264, 2.5264647388870216, 2.5442296421824877],

    # harness
    [2.1037058294469366, 2.1654937099260874, 2.316679448389115, 2.5593173073621953, 2.6884393528150654],

    # kit
    [2.4161212160394556, 2.546064347492432, 2.658053513878462, 2.900830584552272, 2.969184571676013],

    # logrus
    [2.7212449981106257, 2.7544673906194013, 2.7932470475658855, 3.0125727426382145, 3.051819975371107],

    # nsq
    [2.1397274136960127, 2.2047665254527193, 2.4657598203129996, 3.144422890922449, 3.1367371880729764]
]

avg_stdd_nerr_stats = [
    # cobra
    [2.2504048119640294, 2.333960511737758, 2.589203199874845, 2.6820608058278657, 2.7118503589885057],

    # echo
    [2.506401574780886, 2.5962413836889002, 2.591013257562645, 2.727683134324671, 2.9672396108861148],

    # dive
    [2.4541822384254037, 2.432593377648734, 2.6631447684809926, 2.8886123944044657, 3.026046498581449],

    # esbuild
    [1.931387464739194, 2.0084828100250456, 2.2971235887828496, 2.5588465585355866, 2.694292683645554],

    # fiber
    [2.2120354143254866, 2.6194520032104753, 2.7396919942497204, 3.1348747404597805, 3.287589034060858],

    # fzf
    [2.0566105925110154, 2.211601873453724, 2.2226431311748396, 2.6439280377516003, 2.745978218100295],

    # gin
    [1.4456333130241645, 1.6940948779052263, 1.9151766408370121, 2.344177690575173, 2.393480784957241],

    # gorm
    [2.0012345457664202, 2.207973303736503, 2.2844653284085967, 2.5301243230768504, 2.557650579467269],

    # harness
    [2.10060347126852, 2.1590848091653365, 2.314745099993615, 2.5581942665470447, 2.6916092530614013],

    # kit
    [2.4120418404733, 2.543059047090991, 2.6567801512091527, 2.901842836064546, 2.9655000247549794],

    # logrus
    [2.6833359749892427, 2.720637143897362, 2.758364881696795, 2.9669008828796866, 3.0123991095300067],

    # nsq
    [2.153019137063752, 2.218844289556062, 2.4779730990344464, 3.1775401068038898, 3.1698905984669223]
]

avg_cov_stats = [
    # cobra
    [1.8, 1.8, 1.8, 1.8, 1.8],

    # echo
    [12.88, 12.906, 12.906, 12.94, 12.97],

    # dive
    [8.0, 8.0, 8.0, 8.0, 8.0],

    # esbuild
    [2.51, 2.83, 3.43, 4.85, 8.63],

    # fiber
    [8.83, 9.99, 10.29, 10.4, 10.42],

    # fzf
    [3.2, 3.62, 3.98, 4.1, 4.1],

    # gin
    [13.28, 14.04, 14.98, 15.58, 15.79],

    # gorm
    [1.4, 1.4, 1.4, 1.4, 1.4],

    # harness
    [3.68, 5.64, 6.86, 8.9, 8.9],

    # kit
    [5.98, 6.05, 6.07, 6.15, 6.209],

    # logrus
    [29.31, 29.40, 29.42, 29.43, 29.48],

    # nsq
    [11.41, 14.61, 16.95, 18.30, 18.58]
]

def plot_main_stats():
    for i, repo in enumerate(repo_names):
        x  = experiment_times
        y1 = avg_lenght_err_stats[i]
        y2 = avg_lenght_nerr_stats[i]
        y3 = avg_stdd_err_stats[i]
        y4 = avg_stdd_nerr_stats[i]

        plt.figure(figsize=(10, 6))
        plt.scatter(x, y1, label="Tamanho médio", color="red")
        plt.plot(x, y1, color="red")
        plt.scatter(x, y3, label="Desvio padrão", color="green")
        plt.plot(x, y3, color="green")

        plt.title(f"Evolução de {repo} (error_sequences)")
        plt.xlabel("Tempo de experimento (s)")
        plt.ylabel("Métricas")
        plt.xticks(x)
        plt.legend()
        plt.savefig(f"{repo}_evolution_error")
        plt.close()

        plt.figure(figsize=(10, 6))
        plt.scatter(x, y2, label="Tamanho médio", color="red")
        plt.plot(x, y2, color="red")
        plt.scatter(x, y4, label="Desvio padrão", color="green")
        plt.plot(x, y4, color="green")

        plt.title(f"Evolução de {repo} (non_error_sequences)")
        plt.xlabel("Tempo de experimento (s)")
        plt.ylabel("Métricas")
        plt.xticks(x)
        plt.legend()
        plt.savefig(f"{repo}_evolution_nerror")
        plt.close()

def plot_avg_cov_evolution():
    for i, repo in enumerate(repo_names):
        x = experiment_times
        y = avg_cov_stats[i]

        plt.figure(figsize=(10, 6))
        plt.scatter(x, y, label="Coverage médio", color="orange")
        plt.plot(x, y, color="orange")
        plt.title(f"Evolução de coverage - {repo}")
        plt.xlabel("Tempo de experimento (s)")
        plt.ylabel("Coverage médio")
        plt.xticks(x)
        plt.legend()
        plt.savefig(f"{repo}_cov_evolution")

def plot_all_coverages():
    colors = [
        "#4E79A7", "#F28E2B", "#59A14F", "#E15759",
        "#B07AA1", "#9C755F", "#BAB0AC", "#1F77B4",
        "#FFD700", "#17BECF", "#D62728", "#9467BD"
    ]

    markers = ['o', 's', 'D', '^', 'v', 'x', '*', '+', 'p', 'h', '<', '>']

    x = experiment_times
    cobra   = avg_cov_stats[0]
    dive    = avg_cov_stats[1]
    echo    = avg_cov_stats[2]
    esbuild = avg_cov_stats[3]
    fiber   = avg_cov_stats[4]
    fzf     = avg_cov_stats[5]
    gin     = avg_cov_stats[6]
    gorm    = avg_cov_stats[7]
    harness = avg_cov_stats[8]
    kit     = avg_cov_stats[9]
    logrus  = avg_cov_stats[10]
    nsq     = avg_cov_stats[11]

    plt.figure(figsize=(10, 6))
    # cobra
    plt.scatter(x, cobra, label="cobra", color=colors[0])
    plt.plot(x, cobra, color=colors[0])
    # dive
    plt.scatter(x, dive, label="dive", color=colors[1])
    plt.plot(x,    dive,               color=colors[1], marker=markers[1])
    plt.scatter(x, echo, label="echo", color=colors[2], )
    plt.plot(x,    echo,               color=colors[2], marker=markers[2])
    plt.scatter(x, esbuild, label="esbuild", color=colors[3])
    plt.plot(x,    esbuild,               color=colors[3], marker=markers[3])
    plt.scatter(x, fiber, label="fiber", color=colors[4])
    plt.plot(x,    fiber,               color=colors[4], marker=markers[4])
    plt.scatter(x, fzf, label="fzf",  color=colors[5])
    plt.plot(x,    fzf,               color=colors[5], marker=markers[5])
    plt.scatter(x, gin, label="gin", color=colors[6])
    plt.plot(x,    gin,               color=colors[6], marker=markers[6])
    plt.scatter(x, gorm, label="gorm", color=colors[7])
    plt.plot(x,    gorm,               color=colors[7], marker=markers[7])
    plt.scatter(x, harness, label="harness", color=colors[8])
    plt.plot(x,    harness,               color=colors[8], marker=markers[8])
    plt.scatter(x, kit, label="kit", color=colors[9])
    plt.plot(x,    kit,               color=colors[9], marker=markers[9])
    plt.scatter(x, logrus, label="logrus", color=colors[10])
    plt.plot(x,    logrus,               color=colors[10], marker=markers[10])
    plt.scatter(x, nsq, label="nsq", color=colors[11])
    plt.plot(x,    nsq,               color=colors[11], marker=markers[11])

    plt.title(f"Evolução de coverage")
    plt.xlabel("Tempo de experimento (s)")
    plt.ylabel("Coverage médio")
    plt.xticks(x)
    plt.legend()
    plt.savefig(f"all_cov_evolution")

def plot_cov_subplots():
    colors = [
        "#4E79A7", "#F28E2B", "#59A14F", "#E15759",
        "#B07AA1", "#9C755F", "#BAB0AC", "#1F77B4",
        "#FFD700", "#17BECF", "#D62728", "#9467BD"
    ]
    labels = ['(a)', '(b)', '(c)', '(d)']

    x = experiment_times
    cobra   = avg_cov_stats[0]
    dive    = avg_cov_stats[1]
    echo    = avg_cov_stats[2]
    esbuild = avg_cov_stats[3]
    fiber   = avg_cov_stats[4]
    fzf     = avg_cov_stats[5]
    gin     = avg_cov_stats[6]
    gorm    = avg_cov_stats[7]
    harness = avg_cov_stats[8]
    kit     = avg_cov_stats[9]
    logrus  = avg_cov_stats[10]
    nsq     = avg_cov_stats[11]


    # Coverage constante
    plt.figure(figsize=(10,6))
    plt.scatter(x, cobra, label="cobra", color=colors[5])
    plt.plot(x, cobra, color=colors[5])
    plt.scatter(x, dive, label="dive", color=colors[1])
    plt.plot(x,    dive,               color=colors[1])
    plt.scatter(x, gorm, label="gorm",     color=colors[7])
    plt.plot(x,    gorm,                   color=colors[7])
    plt.xticks(x)
    plt.xticks(rotation=45)
    plt.yticks(cobra + dive + gorm)
    plt.tick_params(axis='both', labelsize=12)
    plt.gca().yaxis.set_major_locator(MaxNLocator(nbins=6))
    plt.legend(fontsize=12)
    plt.savefig("subplot_a")
    plt.close()

    # estabilizacao
    plt.figure(figsize=(10,6))
    plt.scatter(x, harness, label="harness", color=colors[8])
    plt.plot(x,    harness,                  color=colors[8])
    plt.scatter(x, fzf, label="fzf",         color=colors[5])
    plt.plot(x,    fzf,                      color=colors[5])
    plt.scatter(x, gin, label="gin",         color=colors[6])
    plt.plot(x,    gin,                      color=colors[6])
    plt.scatter(x, nsq, label="nsq",         color=colors[11])
    plt.plot(x,    nsq,                      color=colors[11])
    plt.yticks(harness + fzf + gin + nsq)
    plt.xticks(x)
    plt.xticks(rotation=45)
    plt.gca().yaxis.set_major_locator(MaxNLocator(nbins=6))
    plt.tick_params(axis='both', labelsize=12)
    plt.legend(fontsize=12)
    plt.savefig("subplot_b")
    plt.close()

    # aumento discreto
    plt.figure(figsize=(10,6))
    plt.scatter(x, echo, label="echo",     color=colors[2])
    plt.plot(x,    echo,                   color=colors[2])
    plt.scatter(x, fiber, label="fiber",   color=colors[4])
    plt.plot(x,    fiber,                  color=colors[4])
    plt.scatter(x, kit, label="kit",       color=colors[9])
    plt.plot(x,    kit,                    color=colors[9])
    plt.scatter(x, logrus, label="logrus", color=colors[10])
    plt.plot(x,    logrus,                 color=colors[10])
    plt.yticks(echo + fiber + kit + logrus)
    plt.tick_params(axis='both', labelsize=12)
    plt.gca().yaxis.set_major_locator(MaxNLocator(nbins=6))
    plt.legend(fontsize=12)
    plt.xticks(x)
    plt.xticks(rotation=45)
    plt.savefig("subplot_c")
    plt.close()

    # aumento contínuo
    plt.figure(figsize=(10, 6))
    plt.scatter(x, esbuild, label="esbuild", color=colors[3])
    plt.plot(x,    esbuild,                  color=colors[3])
    plt.yticks(esbuild)
    plt.xticks(x)
    plt.xticks(rotation=45)
    plt.tick_params(axis='both', labelsize=12)
    plt.legend(fontsize=12)
    # axs[1, 1].yaxis.set_major_locator(MaxNLocator(nbins=6))
    plt.savefig("subplot_d")
    plt.close()


if __name__ == "__main__":
    plot_main_stats()
    plot_avg_cov_evolution()
    plot_all_coverages()
    # plot_cov_subplots()