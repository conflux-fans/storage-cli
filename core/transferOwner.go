package core

// func TransferOwner(name string, from common.Address, to common.Address) error {
// 	// get all keys
// 	logger.Get().WithField("name", name).WithField("from", from).WithField("to", to).Info("Start transfer content owner")

// 	meta, err := GetContentMetadata(name)
// 	if err != nil {
// 		return err
// 	}

// 	keys := append(meta.LineKeys, meta.LineSizeKey)

// 	logger.Get().WithField("length", len(keys)).Info("Get content related keys")

// 	// check is all writer, if not
// 	for _, k := range keys {
// 		isWriter, err := kvClientForIterator.IsWriterOfKey(defaultAccount, STREAM_FILE, []byte(k))
// 		if err != nil {
// 			return errors.WithMessage(err, "Failed to check if owner")
// 		}
// 		if !isWriter {
// 			return fmt.Errorf("not the writer of key %s", k)
// 		}
// 	}

// 	batcher := kvClientsForPut[from].Batcher()

// 	for _, k := range keys {
// 		batcher.GrantSpecialWriteRole(STREAM_FILE, []byte(k), to)
// 		batcher.RenounceSpecialWriteRole(STREAM_FILE, []byte(k))
// 	}

// 	return batcher.Exec()
// }
