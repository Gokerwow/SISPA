import type { Customer } from "@/types/customer";
import apiClient from "./api";

export async function GetAllCustomers() {
  const response = await apiClient.get('/customers')
  return response.data
}

export async function GetCustomerDetail(id: number) {
  const response = await apiClient.get(`/customers/${id}`)
  return response.data
}

export async function AddCustomers(payload: Customer) {
  const response = await apiClient.post('/customers', payload)
  return response.data
}

export async function updateCustomer(id: number, payload: Customer) {
  const response = await apiClient.put(`/customers/${id}`, payload)
  return response.data
}

export async function DeleteCustomers(id: number) {
  const response = await apiClient.delete(`/customers/${id}`)
  return response.data
}
