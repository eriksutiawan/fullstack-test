<template>
    <div class="form-container">
      <form @submit.prevent="handleSubmit">
        <div>
          <label for="name">Name:</label>
          <input type="text" v-model="name" required />
        </div>
        <div>
          <label for="identity_number">Identity Number:</label>
          <input type="text" v-model="identityNumber" required />
        </div>
        <div>
          <label for="email">Email Address:</label>
          <input type="email" v-model="email" required />
        </div>
        <div>
          <label for="date_of_birth">Date of Birth:</label>
          <input type="date" v-model="dateOfBirth" required />
        </div>
        <button type="submit">Submit</button>
      </form>
      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import { useFormStore } from '../stores/form';
  import axios from "axios";
  
  export default defineComponent({
    name: 'RegisterForm',
    setup() {
      const formStore = useFormStore();
  
      const handleSubmit = async () => {
        try {
          const response = await axios.post('http://localhost:8080/api/form', {
            name: formStore.name,
            identity_number: formStore.identityNumber,
            email: formStore.email,
            date_of_birth: formStore.dateOfBirth
          });
          console.log('Form submitted successfully:', response.data);
          formStore.resetForm();
        } catch (error) {
          console.error('Error submitting form:', error);
        }
      };
  
      return {
        name: formStore.name,
        identityNumber: formStore.identityNumber,
        email: formStore.email,
        dateOfBirth: formStore.dateOfBirth,
        errorMessage: '',
        handleSubmit: handleSubmit
      };
    }
  });
  </script>


<style scoped>
.form-container {
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="text"],
input[type="email"],
input[type="date"] {
  width: 100%;
  padding: 8px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button[type="submit"] {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button[type="submit"]:hover {
  background-color: #45a049;
}

button[type="submit"] {
  margin-top: 20px;
}

.error-message {
  color: red;
  margin-top: 10px;
}
</style>
  