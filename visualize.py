import numpy as np
import matplotlib.pyplot as plt
from p4e import generate_primes


def multiplicative_waves(limit):
    """
    Generate a 2D matrix showing multiplicative coverage.

    Rows    = emitter (prime)
    Columns = natural numbers 2..limit
    Value   = 1 if the emitter hits that number, else 0
    """
    primes = generate_primes(limit)
    matrix = np.zeros((len(primes), limit + 1), dtype=int)

    for i, p in enumerate(primes):
        for multiple in range(2 * p, limit + 1, p):
            matrix[i, multiple] = 1

    return primes, matrix


def plot_growth_and_gaps(limit=300):
    """
    Visualise:
    - multiplicative coverage (heatmap of emitter waves)
    - prime distribution (highlighted gaps)
    - damping of gaps with growth density
    """
    primes, waves = multiplicative_waves(limit)

    # Display the coverage matrix
    plt.figure(figsize=(12, 6))
    plt.imshow(
        waves,
        aspect="auto",
        cmap="inferno",
        interpolation="nearest",
        extent=[2, limit, len(primes), 1],
    )
    plt.colorbar(label="multiplicative coverage")
    plt.title("Multiplicative reconstruction waves of the number space")
    plt.xlabel("n")
    plt.ylabel("prime emitters")
    plt.tight_layout()
    plt.show()

    # Plot primes across n
    plt.figure(figsize=(12, 4))
    x = np.arange(2, limit + 1)
    is_prime = np.zeros(limit + 1)
    for p in primes:
        is_prime[p] = 1

    plt.scatter(x, is_prime, s=10)
    plt.title("Prime positions (structural growth gaps)")
    plt.yticks([])
    plt.xlabel("n")
    plt.tight_layout()
    plt.show()

    # Prime gap size growth (damping curve)
    gaps = np.diff(primes)
    plt.figure(figsize=(12, 4))
    plt.plot(primes[1:], gaps, linewidth=1)
    plt.title("Prime gap size over n — damping with growth density")
    plt.xlabel("prime")
    plt.ylabel("gap to next prime")
    plt.tight_layout()
    plt.show()


if __name__ == "__main__":
    plot_growth_and_gaps(500)