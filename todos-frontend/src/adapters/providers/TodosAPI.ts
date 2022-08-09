import { SelectTodosQuery, SelectTodosQueryResult } from '../../domain/messages/SelectTodosQuery';
import { AddTodoCommand } from '../../domain/messages/AddTodoCommand';
import { ClearCompletedCommand } from '../../domain/messages/ClearCompletedCommand';
import { CommandStatus } from '../../domain/messages/CommandStatus';
import { DestroyTodoCommand } from '../../domain/messages/DestroyTodoCommand';
import { SaveTodoCommand } from '../../domain/messages/SaveTodoCommand';
import { ToggleAllCommand } from '../../domain/messages/ToggleAllCommand';
import { ToggleTodoCommand } from '../../domain/messages/ToggleTodoCommand';

import APIUtils from './APIUtils';

const baseUrl = '/api/todos';

async function addTodo(c: AddTodoCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/add-todo`, c);
}

async function clearCompleted(c: ClearCompletedCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/clear-completed`, c);
}

async function destroyTodo(c: DestroyTodoCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/destroy-todo`, c);
}

async function saveTodo(c: SaveTodoCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/save-todo`, c);
}

async function selectTodos(q: SelectTodosQuery): Promise<SelectTodosQueryResult> {
  return APIUtils.getJson(`${baseUrl}/select-todos`, q);
}

async function toggleAll(c: ToggleAllCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/toggle-all`, c);
}

async function toggleTodo(c: ToggleTodoCommand): Promise<CommandStatus> {
  return APIUtils.postJson(`${baseUrl}/toggle-todo`, c);
}

const TodosAPI = {
  addTodo,
  clearCompleted,
  destroyTodo,
  saveTodo,
  selectTodos,
  toggleAll,
  toggleTodo,
};

export default TodosAPI;
