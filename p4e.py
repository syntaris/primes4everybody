def generate_primes(limit: int):
    """
    Constructive prime generator based on the growth model.

    Idea:
    - Iterate through the natural numbers n = 2, 3, 4, ...
    - For every previously discovered prime p we create an "emitter":
         emitter['step'] = p
         emitter['next'] = the next multiple of p not yet reached
    - If the current number n is hit by any emitter → composite
    - If n is not reached by any emitter → prime, and it becomes a new emitter itself

    IMPORTANT:
    - No modulus operator, no divisibility tests, no factorisation.
    - A prime emerges exactly when the existing multiplicative structure cannot recreate n.
    """

    primes = []
    emitters = []

    for n in range(2, limit + 1):
        is_composite = False

        # Check if any emitter generates n as its next multiple
        for emitter in emitters:
            while emitter["next"] < n:
                emitter["next"] += emitter["step"]
            if emitter["next"] == n:
                is_composite = True

        # n is prime if no emitter reaches it
        if not is_composite:
            primes.append(n)
            emitters.append({"next": 2 * n, "step": n})

    return primes


if __name__ == "__main__":
    # Print all primes up to a chosen limit
    for p in generate_primes(20000):
        print(p)