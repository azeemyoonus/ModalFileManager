// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "main": {
    "App": {
      /**
       * AddWatcher
       * @param {string} arg1 - Go Type: string
       * @param {number} arg2 - Go Type: int
       * @param {string} arg3 - Go Type: string
       * @returns {Promise<void>} 
       */
      "AddWatcher": (arg1, arg2, arg3) => {
        return window.go.main.App.AddWatcher(arg1, arg2, arg3);
      },
      /**
       * AppendPath
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "AppendPath": (arg1, arg2) => {
        return window.go.main.App.AppendPath(arg1, arg2);
      },
      /**
       * CloseLeftWatch
       * @returns {Promise<void>} 
       */
      "CloseLeftWatch": () => {
        return window.go.main.App.CloseLeftWatch();
      },
      /**
       * CloseRightWatch
       * @returns {Promise<void>} 
       */
      "CloseRightWatch": () => {
        return window.go.main.App.CloseRightWatch();
      },
      /**
       * CopyEntries
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<void>} 
       */
      "CopyEntries": (arg1, arg2) => {
        return window.go.main.App.CopyEntries(arg1, arg2);
      },
      /**
       * DeleteEntries
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "DeleteEntries": (arg1) => {
        return window.go.main.App.DeleteEntries(arg1);
      },
      /**
       * DirExists
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<boolean>}  - Go Type: bool
       */
      "DirExists": (arg1) => {
        return window.go.main.App.DirExists(arg1);
      },
      /**
       * FileExists
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<boolean>}  - Go Type: bool
       */
      "FileExists": (arg1) => {
        return window.go.main.App.FileExists(arg1);
      },
      /**
       * GetClip
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetClip": () => {
        return window.go.main.App.GetClip();
      },
      /**
       * GetCommandLineCommands
       * @returns {Promise<Array<string>>}  - Go Type: []string
       */
      "GetCommandLineCommands": () => {
        return window.go.main.App.GetCommandLineCommands();
      },
      /**
       * GetEnvironment
       * @returns {Promise<Array<string>>}  - Go Type: []string
       */
      "GetEnvironment": () => {
        return window.go.main.App.GetEnvironment();
      },
      /**
       * GetError
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetError": () => {
        return window.go.main.App.GetError();
      },
      /**
       * GetHomeDir
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetHomeDir": () => {
        return window.go.main.App.GetHomeDir();
      },
      /**
       * MakeDir
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "MakeDir": (arg1) => {
        return window.go.main.App.MakeDir(arg1);
      },
      /**
       * MakeFile
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "MakeFile": (arg1) => {
        return window.go.main.App.MakeFile(arg1);
      },
      /**
       * MoveEntries
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<void>} 
       */
      "MoveEntries": (arg1, arg2) => {
        return window.go.main.App.MoveEntries(arg1, arg2);
      },
      /**
       * Quit
       * @returns {Promise<void>} 
       */
      "Quit": () => {
        return window.go.main.App.Quit();
      },
      /**
       * ReadDir
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<Array<models.FileInfo>>}  - Go Type: []main.FileInfo
       */
      "ReadDir": (arg1) => {
        return window.go.main.App.ReadDir(arg1);
      },
      /**
       * ReadFile
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "ReadFile": (arg1) => {
        return window.go.main.App.ReadFile(arg1);
      },
      /**
       * RemoveWatcher
       * @param {string} arg1 - Go Type: string
       * @param {number} arg2 - Go Type: int
       * @returns {Promise<void>} 
       */
      "RemoveWatcher": (arg1, arg2) => {
        return window.go.main.App.RemoveWatcher(arg1, arg2);
      },
      /**
       * RenameEntry
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<void>} 
       */
      "RenameEntry": (arg1, arg2) => {
        return window.go.main.App.RenameEntry(arg1, arg2);
      },
      /**
       * RunCommandLine
       * @param {string} arg1 - Go Type: string
       * @param {Array<string>} arg2 - Go Type: []string
       * @param {Array<string>} arg3 - Go Type: []string
       * @param {string} arg4 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "RunCommandLine": (arg1, arg2, arg3, arg4) => {
        return window.go.main.App.RunCommandLine(arg1, arg2, arg3, arg4);
      },
      /**
       * SetClip
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "SetClip": (arg1) => {
        return window.go.main.App.SetClip(arg1);
      },
      /**
       * SetLeftDirWatch
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "SetLeftDirWatch": (arg1) => {
        return window.go.main.App.SetLeftDirWatch(arg1);
      },
      /**
       * SetRightDirWatch
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<void>} 
       */
      "SetRightDirWatch": (arg1) => {
        return window.go.main.App.SetRightDirWatch(arg1);
      },
      /**
       * SplitFile
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<models.FileParts>}  - Go Type: main.FileParts
       */
      "SplitFile": (arg1) => {
        return window.go.main.App.SplitFile(arg1);
      },
      /**
       * WriteFile
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<void>} 
       */
      "WriteFile": (arg1, arg2) => {
        return window.go.main.App.WriteFile(arg1, arg2);
      },
    },
  },

};
export default go;
