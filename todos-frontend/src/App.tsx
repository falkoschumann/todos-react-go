import { Route, Routes } from 'react-router-dom';
import { useState } from 'react';

import { SelectTodosQuery, SelectTodosQueryResult } from './domain/messages/SelectTodosQuery';
import { AddTodoCommand } from './domain/messages/AddTodoCommand';
import { ClearCompletedCommand } from './domain/messages/ClearCompletedCommand';
import { DestroyTodoCommand } from './domain/messages/DestroyTodoCommand';
import { SaveTodoCommand } from './domain/messages/SaveTodoCommand';
import { ToggleAllCommand } from './domain/messages/ToggleAllCommand';
import { ToggleTodoCommand } from './domain/messages/ToggleTodoCommand';

import TodosController from './adapters/portals/TodosController';

import TodosAPI from './adapters/providers/TodosAPI';

function App() {
  const [selectedTodos, setSelectedTodos] = useState<SelectTodosQueryResult>();

  async function handleAddTodo(command: AddTodoCommand) {
    const status = await TodosAPI.addTodo(command);
    if (!status.success) {
      console.error('Add todo failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  async function handleClearCompleted(command: ClearCompletedCommand) {
    const status = await TodosAPI.clearCompleted(command);
    if (!status.success) {
      console.log('Clear completed failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  async function handleDestroyTodo(command: DestroyTodoCommand) {
    const status = await TodosAPI.destroyTodo(command);
    if (!status.success) {
      console.log('Destroy todo failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  async function handleSaveTodo(command: SaveTodoCommand) {
    const status = await TodosAPI.saveTodo(command);
    if (!status.success) {
      console.error('Save todo failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  async function handleSelectTodos(query: SelectTodosQuery) {
    const result = await TodosAPI.selectTodos(query);
    setSelectedTodos(result);
  }

  async function handleToggleAll(command: ToggleAllCommand) {
    const status = await TodosAPI.toggleAll(command);
    if (!status.success) {
      console.error('Toggle all failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  async function handleToggleTodo(command: ToggleTodoCommand) {
    const status = await TodosAPI.toggleTodo(command);
    if (!status.success) {
      console.error('Toggle todo failed:', status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }

  return (
    <Routes>
      <Route
        path="/*"
        element={
          <TodosController
            selectedTodos={selectedTodos}
            onAddTodo={handleAddTodo}
            onClearCompleted={handleClearCompleted}
            onDestroyTodo={handleDestroyTodo}
            onSaveTodo={handleSaveTodo}
            onSelectTodos={handleSelectTodos}
            onToggleAll={handleToggleAll}
            onToggleTodo={handleToggleTodo}
          />
        }
      />
    </Routes>
  );
}

export default App;
