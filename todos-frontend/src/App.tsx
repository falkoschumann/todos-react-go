import { Route, Routes } from 'react-router-dom';
import { useCallback, useState } from 'react';

import {
  AddTodoCommand,
  ClearCompletedCommand,
  DestroyTodoCommand,
  SaveTodoCommand,
  SelectTodosQueryResult,
  ToggleAllCommand,
  ToggleTodoCommand,
} from './domain/messages';

import TodosAPI from './adapters/providers/TodosAPI';
import TodosController from './adapters/portals/TodosController';

function App() {
  const [selectedTodos, setSelectedTodos] = useState<SelectTodosQueryResult>();

  const handleAddTodo = useCallback(async (c: AddTodoCommand) => {
    const status = await TodosAPI.addTodo(c);
    if (!status.success) {
      console.error(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleClearCompleted = useCallback(async (c: ClearCompletedCommand) => {
    const status = await TodosAPI.clearCompleted(c);
    if (!status.success) {
      console.log(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleDestroyTodo = useCallback(async (c: DestroyTodoCommand) => {
    const status = await TodosAPI.destroyTodo(c);
    if (!status.success) {
      console.log(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleSaveTodo = useCallback(async (c: SaveTodoCommand) => {
    const status = await TodosAPI.saveTodo(c);
    if (!status.success) {
      console.error(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleSelectTodos = useCallback(async () => {
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleToggleAll = useCallback(async (c: ToggleAllCommand) => {
    const status = await TodosAPI.toggleAll(c);
    if (!status.success) {
      console.error(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

  const handleToggleTodo = useCallback(async (c: ToggleTodoCommand) => {
    const status = await TodosAPI.toggleTodo(c);
    if (!status.success) {
      console.error(status.errorMessage);
    }
    const result = await TodosAPI.selectTodos({});
    setSelectedTodos(result);
  }, []);

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
