import os
import configparser
import logging

from infra_terraform import settings
from infra_terraform.utils import load_configs

# Set up logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s [%(levelname)s] %(message)s',
    handlers=[logging.FileHandler('terraform.log'), logging.StreamHandler()]
)

def main():
    # Load project settings
    config = load_configs()

    # Create Terraform working directory
    terraform_dir = os.path.join(os.path.dirname(__file__), 'terraform')
    if not os.path.exists(terraform_dir):
        os.makedirs(terraform_dir)

    # Create Terraform configuration files
    for resource in config['resources']:
        with open(os.path.join(terraform_dir, f'{resource}.tf'), 'w') as f:
            f.write(load_configs().get_resource(resource))

    # Initialize Terraform
    os.system(f'cd {terraform_dir} && terraform init')

    # Apply Terraform configuration
    os.system(f'cd {terraform_dir} && terraform apply')

if __name__ == '__main__':
    main()