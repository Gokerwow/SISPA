import apiClient from "./api";
import type { TransactionRequest, UpdateStatusRequest } from "@/types/transaction";

export async function GetAllHeader() {
  const response = await apiClient.get('/transactions')
  return response.data
}

export async function GetInvoiceDetails(id: number) {
  const response = await apiClient.get(`/transactions/${id}`)
  return response.data
}

export async function AddTransaction(payload: TransactionRequest) {
  const response = await apiClient.post('/transactions', payload)
  return response.data
}

export async function updateStatusTransaction(id: number, payload: UpdateStatusRequest) {
  const response = await apiClient.patch(`/transactions/${id}`, payload)
  return response.data
}

export async function updateTransaction(id: number, payload: TransactionRequest) {
  const response = await apiClient.put(`/transactions/${id}`, payload)
  return response.data
}
