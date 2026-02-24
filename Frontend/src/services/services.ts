import type { Services, ServicesCreate } from "@/types/services";
import apiClient from "./api";

export async function GetAllServices() {
  const response = await apiClient.get('/services')
  return response.data
}

export async function GetServiceDetail(id: number) {
  const response = await apiClient.get(`/services/${id}`)
  return response.data
}

export async function AddServices(payload: ServicesCreate) {
  const response = await apiClient.post('/services', payload)
  return response.data
}

export async function updateService(id: number, payload: Services) {
  const response = await apiClient.put(`/services/${id}`, payload)
  return response.data
}

export async function DeleteService(id: number) {
  const response = await apiClient.delete(`/services/${id}`)
  return response.data
}
