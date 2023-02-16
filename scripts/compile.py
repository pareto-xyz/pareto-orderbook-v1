"""
Run this file to compile the MarginV1 smart contract to a Go file. 
See https://github.com/pareto-xyz/pareto-core-v1. 

Requires `solc` and `abigen` to be globally callable. 

To install `solc`, see https://docs.soliditylang.org/en/v0.8.9/installing-solidity.html.
To install `abigen`, see https://geth.ethereum.org/docs/dapp/native-bindings. 
"""

from os import makedirs
from os.path import join, basename, splitext
import subprocess


def main(args):
    r"""Compiles contracts and generates go scripts"""
    compile_contract(out_name='margin',
                     contract_path=join(args.core_dir, 'contracts/MarginV1.sol'), 
                     out_dir=args.out_dir,
                     dependencies={
                        # NOTE: the final "/" is very important
                        '@openzeppelin/': join(args.core_dir, 'node_modules/@openzeppelin/')
                     })
    compile_contract(out_name='oracle',
                     contract_path=join(args.core_dir, 'contracts/oracles/Oracle.sol'), 
                     out_dir=args.out_dir,
                     dependencies={
                        # NOTE: the final "/" is very important
                        '@openzeppelin/': join(args.core_dir, 'node_modules/@openzeppelin/'),
                        'IOracle.sol': join(args.core_dir, 'contracts/interfaces/IOracle.sol'),
                     })


def compile_contract(out_name, contract_path, out_dir, dependencies={}):
    makedirs(out_dir, exist_ok=True)
    name = splitext(basename(contract_path))[0]

    abi_dir = join(out_dir, out_name, 'abi')
    bin_dir = join(out_dir, out_name, 'bin')
    makedirs(abi_dir, exist_ok=True)
    makedirs(bin_dir, exist_ok=True)

    extra_args = ' '.join([f'{k}={v}' for k, v in dependencies.items()])
    abi_cmd = f"solc --abi {extra_args} {contract_path} -o {abi_dir} --overwrite"
    bin_cmd = f"solc --bin {extra_args} {contract_path} -o {bin_dir} --overwrite"

    print(f"Running: {abi_cmd}")
    run_bash(abi_cmd)
    
    print(f"Running: {bin_cmd}")
    run_bash(bin_cmd)

    abigen_args = {
        'bin': join(bin_dir, f'{name}.bin'),
        'abi': join(abi_dir, f'{name}.abi'),
        'pkg': out_name,
        'out': join(out_dir, out_name, f'{out_name}.go'),
    }
    abigen_args = dict_to_args(abigen_args)
    abigen_cmd = f"abigen {abigen_args}"

    print(f"Running: {abigen_cmd}")
    run_bash(abigen_cmd)

    print(f"Compiled {out_name} contract.")


def dict_to_args(x):
    """
    Helper function to convert a dictionary of keys and values 
    to command line settings

    Arguments:
    --
    x: dict[string, integer]

    Example:
    --
    {'a': 1, 'b': 2} => --a=1 --b=2
    """
    out = ''
    for k, v in x.items():
        out += f'--{k}={v}'
        out += ' '
    return out.strip()


def run_bash(cmd):
    """
    Helper function to run a bash command

    Arguments:
    --
    cmd: string
        Command to run in bash
    """
    subprocess.call(cmd, shell=True, stdout=subprocess.PIPE)


if __name__ == "__main__":
    import argparse 
    parser = argparse.ArgumentParser()
    parser.add_argument("core_dir", type=str, help="Path to the pareto-core-v1 repo")
    parser.add_argument("out_dir", type=str, help="Output directory")
    args = parser.parse_args()
    main(args)
