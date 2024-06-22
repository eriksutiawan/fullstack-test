import { shallowMount } from '@vue/test-utils';
import { describe, it, expect } from 'vitest'
import RegisterForm from '@/components/RegisterForm.vue';

describe('RegisterForm.vue', () => {
  it('renders a RegisterForm', () => {
    const wrapper = shallowMount(RegisterForm);
    expect(wrapper.find('registerForm').exists()).toBe(true);
  });
});
