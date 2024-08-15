# **Some Notes**
- The Ansible scripts only work on Debian family Linux distributions.

# **Installation**
Follow the [installation guide](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html).

# **Running the Playbook**
```shell
ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/07-run-traefik.yaml
```