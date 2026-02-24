import type { Therapist, TherapistCreate } from "@/types/therapist";
import apiClient from "./api";

export async function GetAllTherapist() {
  const response = await apiClient.get('/therapists')
  return response.data
}

export async function GetTherapistDetail(id: number) {
  const response = await apiClient.get(`/therapists/${id}`)
  return response.data
}

export async function AddTherapist(payload: TherapistCreate) {
  const response = await apiClient.post('/therapists', payload)
  return response.data
}

export async function updateTherapist(id: number, payload: Therapist) {
  const response = await apiClient.put(`/therapists/${id}`, payload)
  return response.data
}

export async function DeleteTherapist(id: number) {
  const response = await apiClient.delete(`/therapists/${id}`)
  return response.data
}
