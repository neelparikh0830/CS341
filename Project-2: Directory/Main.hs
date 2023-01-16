{-# LANGUAGE NamedFieldPuns #-}
module Main where

import System.Directory (doesFileExist)
import Data.List (insert, intercalate)

fileName = "directory"
main = readDirectory >>= processQueries

readDirectory = do
    exists <- doesFileExist fileName
    if exists then do
        content <- readFile fileName
        entries <- mapM readEntry $ lines content
        putStrLn "file read successfully"
        return entries
    else do
        putStrLn "file doesn't exist yet"
        return []

readEntry line =
    case words line of
        [firstName, lastName, telephoneNumber] ->
            return $ Entry firstName lastName telephoneNumber

data Entry = Entry
    { firstName, lastName, telephoneNumber :: String
    } deriving (Eq, Ord)

instance Show Entry where
    show Entry {firstName, lastName, telephoneNumber} =
        unwords [firstName, lastName, telephoneNumber]

data Query = Exit | List | Add Entry | Remove Int

readQuery line =
    case words line of
        ["x"] ->
            Just Exit
        ["l"] ->
            Just List
        ["a", firstName, lastName, telephoneNumber] ->
            Just $ Add $ Entry firstName lastName telephoneNumber
        ["r", index] ->
            Just $ Remove $ read index

processQueries directory = do
    line <- getLine
    case readQuery line of
        Just Exit -> do
            writeFile fileName (intercalate "\n" $ map show directory)
            putStrLn "exiting"
        Just query ->
            processQuery query directory >>= processQueries

processQuery Exit = return
processQuery List = \entries -> do
    let index i x = show i ++ ": " ++ show x
    case entries of
        first : rest -> do
            putStrLn "entries:"
            putStrLn $ "[ " ++ index 0 first
            mapM_ (putStrLn . (", " ++)) (zipWith index [1..] rest)
            putStrLn "]"

    return entries
processQuery (Add i) = \entries -> do
    putStrLn "entry added successfuly"
    return $ insert i entries
processQuery (Remove i) = \entries ->
    if i >= 0 && i < length entries then do
        let (left, x : right) = splitAt i entries
        putStrLn "entry removed successfuly"
        return $ left ++ right
    else do
        putStrLn "out of bounds"
        return entries