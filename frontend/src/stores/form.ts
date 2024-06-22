import { defineStore } from 'pinia';

interface FormState {
  name: string;
  identityNumber: string;
  email: string;
  dateOfBirth: string;
}

export const useFormStore = defineStore('form', {
  state: (): FormState => ({
    name: '',
    identityNumber: '',
    email: '',
    dateOfBirth: ''
  }),
  actions: {
    setName(name: string) {
      this.name = name;
    },
    setIdentityNumber(identityNumber: string) {
      this.identityNumber = identityNumber;
    },
    setEmail(email: string) {
      this.email = email;
    },
    setDateOfBirth(dateOfBirth: string) {
      this.dateOfBirth = dateOfBirth;
    },
    resetForm() {
      this.name = '';
      this.identityNumber = '';
      this.email = '';
      this.dateOfBirth = '';
    }
  }
});
