import os
import argparse
from infra_terraform.terraform import Terraform

def main():
    parser = argparse.ArgumentParser(description='Terraform Management Tool')
    parser.add_argument('-a', '--action', choices=['init', 'apply', 'destroy', 'plan'], required=True)
    parser.add_argument('-f', '--file', default='main.tf')
    parser.add_argument('-c', '--config', default='config.tfvars')
    args = parser.parse_args()

    if args.action == 'init':
        terraform = Terraform(args.file)
        terraform.init()
    elif args.action == 'apply':
        terraform = Terraform(args.file, args.config)
        terraform.apply()
    elif args.action == 'destroy':
        terraform = Terraform(args.file)
        terraform.destroy()
    elif args.action == 'plan':
        terraform = Terraform(args.file, args.config)
        terraform.plan()

if __name__ == "__main__":
    main()